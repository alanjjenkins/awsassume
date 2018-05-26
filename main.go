package main

import (
	"log"
	"os"

	"github.com/alanjjenkins/awsassume/awsassume"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "awsassume"
	app.Usage = "Assume AWS IAM roles for secure usage of AWS accounts."
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "browser, b",
			Value: "xdg-open",
			Usage: "The browser to use to log into the account.",
		},
		cli.BoolFlag{
			Name: "envvars, e",
		},
	}

	app.Action = func(c *cli.Context) error {
		config := awsassume.ParseConfig()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
