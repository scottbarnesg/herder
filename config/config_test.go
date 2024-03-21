package config

import (
	"log"
	"testing"
)

func TestReadConfig(t *testing.T) {
	configFilePath := "../example.yml"
	config := ReadConfig(configFilePath)
	log.Printf("%+v\n", config)
}
