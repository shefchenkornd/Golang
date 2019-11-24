package main

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	Enabled bool
	Path string
}

func main() {
	fmt.Println("Parsing YAML file")

	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	var yamlConfig YamlConfig
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	fmt.Printf("Result: %v\n", yamlConfig)
	fmt.Println(yamlConfig.Path)
}
