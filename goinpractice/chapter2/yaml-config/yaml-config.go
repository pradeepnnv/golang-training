package main

import (
	"log"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(config)
	log.Println(config.Get("path"))
	log.Println(config.Get("enabled"))
}
