package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"home"
	"io/ioutil"
	"log"
	"net/http"
)

type config struct {
	port   string
	param1 string
	param2 string
}

var defaultConfig config = config{port: "80", param1: "value1", param2: "value2"}

func main() {
	// load config file
	configItems, err := loadConfig()
	if err != nil {
		log.Printf("loadConfig error: %v\n", err)
		configItems = defaultConfig
	}

	// parse args
	configItems, err = parseArgs()
	if err != nil {
		log.Printf("parseArgs error: %v\n", err)
		configItems = defaultConfig
	}

	// register routers
	http.HandleFunc("/", home.HomePage)

	// listen and serve
	addr := fmt.Sprintf(":%s", configItems.port)
	http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("http.ListenAndServe error: %v", err)
	}
}

func loadConfig() (configItems config, err error) {

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &configItems)
	return
}

func parseArgs() (configItems config, err error) {

	err = nil
	flag.StringVar(&configItems.port, "p", "80", "server port")
	flag.Parse()
	return
}
