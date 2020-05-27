package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config represents the YAML config file
type Config struct {
	Global struct {
		Debug                bool
		Port                 uint16
		Domains              []string
		AllowedOrigins       []string
		Directories          []string
		ObfuscateFilenames   bool `yaml:"obfuscate filenames,omitempty"`
		CacheControlLifetime uint32 `yaml:"cache-control lifetime"`
		Logs				 struct {
			Access			 string
			Error			 string	
		}
	}
}

// LoadConfig loads and processes a configuration file
func LoadConfig(filename string) Config {

	// gosec flags the following line as Potential file inclusion via variable 
	// but the filename passed in is used to read to a variable & parsed as YAML
	yamlString, fileErr := ioutil.ReadFile(filename) // #nosec
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	config := Config{}

	err := yaml.Unmarshal([]byte(yamlString), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, o := range config.Global.Domains {
		config.Global.AllowedOrigins = append(config.Global.AllowedOrigins, fmt.Sprintf("http://%s", o))
		config.Global.AllowedOrigins = append(config.Global.AllowedOrigins, fmt.Sprintf("https://%s", o))
	}

	return config
}
