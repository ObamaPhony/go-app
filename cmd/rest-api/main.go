package main

import (
	"flag"
	"fmt"

	config "git.shymega.org.uk/obamaphony/rest-api/internal/config"
	controllers "git.shymega.org.uk/obamaphony/rest-api/internal/controllers"

	log "github.com/inconshreveable/log15"
)

var configPath string

func init() {
	log.Info("Initialising ObamaPhony REST API..")

	/* Setup flags */
	flag.StringVar(&configPath, "configPath",
		"./config.json",
		"Path to configuration file")
	flag.Parse()
}

func main() {
	cfg := config.LoadConfig(configPath)

	bindAddr := fmt.Sprintf("%s:%d",
		cfg.Listener.HTTP.BindAddress,
		cfg.Listener.HTTP.BindPort)

	log.Info("*** ObamaPhony REST API Version 0.1.0 loaded ***")

	go controllers.Server(bindAddr)

	select {}
}
