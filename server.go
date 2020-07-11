package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	chat "github.com/rodaine/grpc-chat/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const tokenHeader = "x-chat-token"

type server struct {
	Host, Password string

	Broadcast chan *chat.StreamResponse

	ClientNames   map[string]string
	ClientStreams map[string]chan *chat.StreamResponse

	namesMtx, streamsMtx sync.RWMutex
}

func Server(host, pass string) *server {
	return &server{
		Host:     host,
		Password: pass,

		Broadcast: make(chan *chat.StreamResponse, 1000),

		ClientNames:   make(map[string]string),
		ClientStreams: make(map[string]chan *chat.StreamResponse),
	}
}

func (s *server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ServerLogf(time.Now(),
		"starting on %s with password %q",
		s.Host, s.Password)

	srv := grpc.NewServer()
	chat.RegisterChatServer(srv, s)

	l, err := net.Listen("tcp", s.Host)
	if err != nil {
		return errors.WithMessage(err,
			"server unable to bind on provided host")
	}

	go s.broadcast(ctx)

	go func() {
		_ = srv.Serve(l)
		cancel()
	}()

	<-ctx.Done()

	s.Broadcast <- &chat.StreamResponse{
		Timestamp: ptypes.TimestampNow(),
		Event: &chat.StreamResponse_ServerShutdown{
			ServerShutdown: &chat.StreamResponse_Shutdown{}}}

	close(s.Broadcast)
	ServerLogf(time.Now(), "shutting down")

	srv.GracefulStop()
	return nil
}

func (s *server) Login(_ context.Context, req *chat.LoginRequest) (*chat.LoginResponse, error) {
	switch {
	case req.Password != s.Password:
		return nil, status.Error(codes.Unauthenticated, "password is incorrect")
	case req.Name == "":
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	tkn := s.genToken()
	s.setName(tkn, req.Name)

	ServerLogf(time.Now(), "%s (%s) has logged in", tkn, req.Name)

	s.Broadcast <- &chat.StreamResponse{
		Timestamp: ptypes.TimestampNow(),
		Event: &chat.StreamResponse_ClientLogin{ClientLogin: &chat.StreamResponse_Login{
			Name: req.Name,
		}},
	}

	return &chat.LoginResponse{Token: tkn}, nil
}

func (s *server) Logout(_ context.Context, req *chat.LogoutRequest) (*chat.LogoutResponse, error) {
	name, ok := s.delName(req.Token)
	if !ok {
		return nil, status.Error(codes.NotFound, "token not found")
	}

	ServerLogf(time.Now(), "%s (%s) has logged out", req.Token, name)

	s.Broadcast <- &chat.StreamResponse{
		Timestamp: ptypes.TimestampNow(),
		Event: &chat.StreamResponse_ClientLogout{ClientLogout: &chat.StreamResponse_Logout{
			Name: name,
		}},
	}

	return new(chat.LogoutResponse), nil
}

func (s *server) Stream(srv chat.Chat_StreamServer) error {
	tkn, ok := s.extractToken(srv.Context())
	if !ok {
		return status.Error(codes.Unauthenticated, "missing token header")
	}

	name, ok := s.getName(tkn)
	if !ok {
		return status.Error(codes.Unauthenticated, "invalid token")
	}

	go s.sendBroadcasts(srv, tkn)

	for {
		req, err := srv.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		s.Broadcast <- &chat.StreamResponse{
			Timestamp: ptypes.TimestampNow(),
			Event: &chat.StreamResponse_ClientMessage{ClientMessage: &chat.StreamResponse_Message{
				Name:    name,
				Message: req.Message,
			}},
		}
	}

	<-srv.Context().Done()
	return srv.Context().Err()
}

func (s *server) sendBroadcasts(srv chat.Chat_StreamServer, tkn string) {
	stream := s.openStream(tkn)
	defer s.closeStream(tkn)

	for {
		select {
		case <-srv.Context().Done():
			return
		case res := <-stream:
			if s, ok := status.FromError(srv.Send(res)); ok {
				switch s.Code() {
				case codes.OK:
					// noop
				case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
					DebugLogf("client (%s) terminated connection", tkn)
					return
				default:
					ClientLogf(time.Now(), "failed to send to client (%s): %v", tkn, s.Err())
					return
				}
			}
		}
	}
}

func (s *server) broadcast(_ context.Context) {
	for res := range s.Broadcast {
		s.streamsMtx.RLock()
		for _, stream := range s.ClientStreams {
			select {
			case stream <- res:
				// noop
			default:
				ServerLogf(time.Now(), "client stream full, dropping message")
			}
		}
		s.streamsMtx.RUnlock()
	}
}

func (s *server) openStream(tkn string) (stream chan *chat.StreamResponse) {
	stream = make(chan *chat.StreamResponse, 100)

	s.streamsMtx.Lock()
	s.ClientStreams[tkn] = stream
	s.streamsMtx.Unlock()

	DebugLogf("opened stream for client %s", tkn)

	return
}

func (s *server) closeStream(tkn string) {
	s.streamsMtx.Lock()

	if stream, ok := s.ClientStreams[tkn]; ok {
		delete(s.ClientStreams, tkn)
		close(stream)
	}

	DebugLogf("closed stream for client %s", tkn)

	s.streamsMtx.Unlock()
}

func (s *server) genToken() string {
	tkn := make([]byte, 4)
	rand.Read(tkn)
	return fmt.Sprintf("%x", tkn)
}

func (s *server) getName(tkn string) (name string, ok bool) {
	s.namesMtx.RLock()
	name, ok = s.ClientNames[tkn]
	s.namesMtx.RUnlock()
	return
}

func (s *server) setName(tkn string, name string) {
	s.namesMtx.Lock()
	s.ClientNames[tkn] = name
	s.namesMtx.Unlock()
}

func (s *server) delName(tkn string) (name string, ok bool) {
	name, ok = s.getName(tkn)

	if ok {
		s.namesMtx.Lock()
		delete(s.ClientNames, tkn)
		s.namesMtx.Unlock()
	}

	return
}

func (s *server) extractToken(ctx context.Context) (tkn string, ok bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md[tokenHeader]) == 0 {
		return "", false
	}

	return md[tokenHeader][0], true
}
