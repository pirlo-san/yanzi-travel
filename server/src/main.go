package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"idebug"
	"ihome"
	"io/ioutil"
	"log"
	"net/http"
)

type config struct {
	Debug bool
}

func main() {
	// load config file
	configItems, err := loadConfig()
	if err != nil {
		log.Printf("loadConfig error: %v\n", err)
	}

	// parse args
	port, err := parsePort()
	if err != nil {
		log.Printf("parsePort error: %v, use default port 80\n", err)
		port = "80"
	}

	log.Println(configItems)
	// set debug mode
	idebug.SetDebugMode(configItems.Debug)

	// register routers
	http.HandleFunc("/", ihome.HomePage)

	// listen and serve
	addr := fmt.Sprintf(":%s", port)
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

func parsePort() (port string, err error) {

	err = nil
	flag.StringVar(&port, "p", "80", "server port")
	flag.Parse()
	return
}
