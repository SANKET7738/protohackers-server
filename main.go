package main

import (
	"flag"
	"log"

	"github.com/SANKET7738/protohackers-server/config"
	"github.com/SANKET7738/protohackers-server/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host server")
	flag.IntVar(&config.Port, "port", 7379, "port for device server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("Starting the server")
	server.RunSyncTCPServer()

}
