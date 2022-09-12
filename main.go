package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	swos "swosman/swos"

	"gopkg.in/yaml.v2"
)

func main() {
	swosClientConfig := swos.Config{}
	swosClientConfig.Host = "172.20.17.213"
	swosClientConfig.HttpProtocol = "http"
	swosClientConfig.Port = 80
	swosClientConfig.Username = "admin"
	swosClientConfig.Password = ""
	client, err := swos.NewClient(&swosClientConfig)
	if err != nil {
		log.Println(err)
		return
	}

	filejson, err := json.MarshalIndent(client.Sw, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	_ = ioutil.WriteFile("./raw/swosInterfaces.json", filejson, 0644)
	fileyaml, err := yaml.Marshal(client.Sw)
	if err != nil {
		log.Println(err)
		return
	}
	_ = ioutil.WriteFile("./raw/swosInterfaces.yaml", fileyaml, 0644)

}
