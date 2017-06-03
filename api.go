package main

import (
	"flag"
	"io/ioutil"
	"os"

	log "github.com/inconshreveable/log15"
	"github.com/jeffail/gabs"
	"github.com/obamaphony/rest-api/controllers"
)

var config *gabs.Container

func init() {
	/* Setup flags */
	configPath := flag.String("configPath", "", "Path to configuration file")
	flag.Parse()

	/* Initialise configuration instance, w/ GABS. */
	log.Debug("Initialise configuration instance")

	/* Ascertain if the configuration path has been inputted. */
	log.Debug("Testing if the configuration path is inputted.")

	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Crit("The configuration file does not exist!",
			log.Ctx{"File": *configPath})
		os.Exit(1)
	}

	log.Debug("Reading configuration file into RAM.")
	configFile, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Error("The JSON config could not be loaded..",
			log.Ctx{"Error": err.Error()})
		os.Exit(1)
	}

	log.Debug("Parsing configuration file into a GABS instance.")
	config, err = gabs.ParseJSON(configFile)
	if err != nil {
		log.Error("The JSON config could not be parsed.",
			log.Ctx{"Error": err.Error()})
		os.Exit(1)
	}

	/* Finish initializing the configuration instance */
}

func main() {
	log.Info("Initialising ObamaPhony REST API..")
	log.Info("*** ObamaPhony REST API Version 0.1.0 loaded ***")

	bindAddr := config.Path("listener.http.bindAddress").Data().(string)

	controllers.Server(bindAddr)
}
