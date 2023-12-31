package main // import "github.com/obamaphony/go-app/cmd/app"

import (
	"flag"
	"fmt"

	config "github.com/obamaphony/go-app/internal/config"
	controllers "github.com/obamaphony/go-app/internal/controllers"

	log "github.com/inconshreveable/log15"
)

var (
	configPath string
	bindAddr   string
)

const defaultConfigPath string = "./config.json"

func init() {
	log.Info("Initialising ObamaPhony REST API..")

	/* Setup flags */
	flag.StringVar(&configPath, "config",
		defaultConfigPath, "Path to configuration file")
	flag.Parse()
}

func main() {
	cfg := config.LoadConfig(configPath)

	bindAddr = fmt.Sprintf("%s:%d",
		cfg.Listener.HTTP.BindAddress,
		cfg.Listener.HTTP.BindPort)

	log.Info("*** ObamaPhony REST API Version 0.1.0 loaded ***")

	go controllers.Server(bindAddr)

	select {}
}
