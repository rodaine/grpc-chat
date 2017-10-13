package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/net/context"
)

var (
	serverMode bool
	debugMode  bool
	host       string
	password   string
	username   string
)

func init() {
	flag.BoolVar(&serverMode, "s", false, "run as the server")
	flag.BoolVar(&debugMode, "v", false, "enable debug logging")
	flag.StringVar(&host, "h", "0.0.0.0:6262", "the chat server's host")
	flag.StringVar(&password, "p", "", "the chat server's password")
	flag.StringVar(&username, "n", "", "the username for the client")
	flag.Parse()
}

func init() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
}

func main() {
	ctx := SignalContext(context.Background())
	var err error

	if serverMode {
		DebugLogf("server mode")
		err = Server(host, password).Run(ctx)
	} else {
		DebugLogf("client mode")
		err = Client(host, password, username).Run(ctx)
	}

	if err != nil {
		MessageLog(time.Now(), "<<Process>>", err.Error())
		os.Exit(1)
	}
}
