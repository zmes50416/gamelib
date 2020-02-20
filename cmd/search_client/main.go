package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zmes50416/gl/internal/kg"
)

func main() {
	app := &cli.App{
		Name:  "search local client",
		Usage: "search Google Knowledge graph",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "api",
				Value: "",
				Usage: "apikey for querying Google Knowledge graph",
			},
		},
		Action: func(c *cli.Context) error {
			key := c.String("api")
			if key != "" {
				kg.APIKey = key
			}
			resp, err := kg.SearchGame("goose")
			if err != nil {
				log.Println(string(resp))
				log.Fatal(err)
			}
			log.Println("Response: ", string(resp))
			return nil
		},
	}
	app.Run(os.Args)
	// := os.Args[1]
}
