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
		Debug              bool
		Port               uint16
		Domains            []string
		AllowedOrigins     []string
		Directories        []string
		ObfuscateFilenames bool `yaml:"obfuscate filenames,omitempty"`
	}
}

// LoadConfig loads and processes a configuration file
func LoadConfig(filename string) Config {

	yamlString, fileErr := ioutil.ReadFile(filename)
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

	fmt.Println(config.Global.ObfuscateFilenames)
	return config
}
