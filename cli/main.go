package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tera-insights/go-aspera"
	"github.com/urfave/cli/v2"
)

func main() {

	// Set log output file
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)

	log.Print("Starting Aspera CLI")
	app := (&cli.App{

		Name:      "aspera",
		Usage:     "Aspera CLI",
		UsageText: "ti-ascli",
		Action: func(c *cli.Context) error {
			log.Print("Running Aspera CLI")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "Authenticate",
				Aliases: []string{"a"},
				Usage:   "Authenticate a user",
				Action: func(c *cli.Context) error {
					log.Print("Authenticating user")
					username := c.Args().Get(0)
					password := c.Args().Get(1)
					if username == "" || password == "" {
						log.Print("Username and password are empty")
						fmt.Println("Username and password are required")
						return nil
					}

					client := aspera.NewClient(&http.Client{}, "demo.asperasoft.com")
					authService := aspera.NewAuthenticateService(client)
					if err := authService.Authenticate(c.Context, &aspera.AuthSpec{
						RemoteUser:     username,
						RemotePassword: password,
					}); err != nil {
						log.Print(err)
						fmt.Printf("Authentication failed: %v\n", err)
					}
					return nil
				},
			},
			{
				Name:    "Upload",
				Aliases: []string{"u"},
				Usage:   "Upload a file",
				Action: func(c *cli.Context) error {
					log.Print("Uploading file")
					return nil
				},
			},
		},
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
