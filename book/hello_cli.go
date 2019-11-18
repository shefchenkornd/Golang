package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	// создание нового экземпляра приложения CLI
	app := cli.App{
		Name: "hello_cli",
		Usage: "Who to say hello to",
		Action: func(c *cli.Context) error {
			// name := c.GlobalString("name")
			name := "John"
			fmt.Printf("Hello, %s\n", name)
			return nil
		},
	}

	// Запуск приложения
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
