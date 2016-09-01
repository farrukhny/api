package common

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Address      string
	Listener     string
	TLSCertFile  string
	TLSKeyFile   string
	ReadTimeout  string
	WriteTimeout string
	MongoDBHost  string
	DBUser       string
	DBPassword   string
	Database     string
}

var AppConfig configuration

func initConfig() {
	loadConfig()
}

func loadConfig() {
	configFile, err := os.Open("config/config.json")
	defer configFile.Close()
	if err != nil {
		log.Fatalf("loadConfig - %s\n", err)
	}

	decoder := json.NewDecoder(configFile)
	AppConfig = configuration{}

	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("decoder - %s\n", err)
	}
}
