package config

import (
	"encoding/json"
	"os"

	log "github.com/inconshreveable/log15"
)

// Config defines the structure for the configuration file (JSON).
type Config struct {
	Listener struct {
		HTTP struct {
			BindAddress string `json:"bindAddress"`
			BindPort    int    `json:"bindPort"`
		} `json:"http"`
	} `json:"listener"`
}

// LoadConfig will load the configuration from `file` argument.
// Should be JSON.
func LoadConfig(path string) Config {
	config := new(Config)

	file, err := os.Open(path)
	if err != nil {
		log.Error("Unable to open configuration file.",
			log.Ctx{"File": path,
				"Error": err.Error()})
		os.Exit(-1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error("Could not defer file close.",
				log.Ctx{"Error": err.Error()})
			os.Exit(-1)
		}
	}(file)

	if err = json.NewDecoder(file).Decode(&config); err != nil {
		log.Error("Error parsing configuration file.",
			log.Ctx{"File": path,
				"Error": err.Error()})
		os.Exit(-1)
	}

	return *config
}
