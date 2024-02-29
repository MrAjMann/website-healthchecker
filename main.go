package main


import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
 )

 func main() {
	app := &cli.App{
		Name: "Healthchecker",
		Usage: "Check the health of a website",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "domain",
				Aliases: []string{"d"},
				Usage: "The domain name to check",
				Required: true,
		},
		&cli.StringFlag{
			Name: "port",
			Aliases: []string{"p"},
			Usage: "The port to check",
			Required: true,
		},
	},

		Action: func(c *cli.Context) error {
			port := c.String("port") 
			if c.String("port") == "" {
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
 }