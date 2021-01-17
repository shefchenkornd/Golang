package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Определяем структуру, соответветствующую значениям в config.json
type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	conf := configuration{}
	errDecode := decoder.Decode(&conf)
	if errDecode != nil {
		fmt.Println("Error:", errDecode)
	}

	fmt.Println(conf.Path)
}
