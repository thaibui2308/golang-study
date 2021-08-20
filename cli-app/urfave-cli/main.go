package main

import (
	"fmt"
	"os"

	"example.com/goinpractice/cli-app/urfave-cli/controller"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Usage = "Count up or down."

	app.Commands = []*cli.Command{
		{
			Name:  "fetch",
			Usage: "Get covid cases in a country.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "country, c",
					Usage: "Specify which country to crawl data",
					Value: "vietnam",
				},
				&cli.StringFlag{
					Name:  "target,t",
					Usage: "Used to target the desired data",
					Value: "todayCases",
				},
			},
			Action: func(c *cli.Context) error {
				country := c.String("country")
				target := c.String("target")
				responseData := controller.GetCountryCovidCase(country)
				if target != "" {
					targetedData := responseData.TodayCases
					fmt.Printf("Today cases: %s", targetedData)
				}
				fmt.Println(responseData)
				return nil
			},
		},
	}

	app.Run(os.Args)

	// app.Usage = "Print welcoming message"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "name, n",
			Value: "World",
			Usage: "Print welcome message.",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.String("name")
		fmt.Printf("Hello %s!\n", name)
		return nil
	}
}
