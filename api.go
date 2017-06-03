package main

import (
	"io/ioutil"
	"os"

	log "github.com/inconshreveable/log15"
	"github.com/jeffail/gabs"
	"github.com/obamaphony/rest-api/controllers"
)

var config *gabs.Container

func init() {
	/* Initialise configuration instance, w/ GABS. */

	log.Debug("Initialise configuration instance")

	// This has to be a absolute path.
	configurationPath := os.Getenv("RESTAPI_CONFIG")

	/* Ascertain if the configuration path is input. */

	log.Debug("Testing if the configuration path is inputted.")
	if _, err := os.Stat(configurationPath); os.IsNotExist(err) {
		log.Crit("The configuration file does not exist!",
			log.Ctx{"File": configurationPath})
		os.Exit(1)
	}

	log.Debug("Reading configuration file into RAM.")
	configurationFile, err := ioutil.ReadFile(configurationPath)
	if err != nil {
		log.Error("The JSON config could not be loaded..",
			log.Ctx{"Error": err.Error()})
		os.Exit(1)
	}

	log.Debug("Parsing configuration file into a GABS instance.")
	config, err = gabs.ParseJSON(configurationFile)
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

	controllers.Server()
}
