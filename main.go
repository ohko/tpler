package main

import (
	"flag"

	"github.com/ohko/tpler/controller"
)

var (
	addr         = flag.String("s", ":8080", "server address")
	sessionPath  = flag.String("sp", "/tmp/hst_session", "session path")
	oauth2Server = flag.String("o2", "https://oauth2.cdeyun.com", "oauth2 server")
)

func main() {
	flag.Parse()

	controller.Start(*addr, *sessionPath, *oauth2Server)
}
