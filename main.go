package main

import (
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cut-the-crap"
	app.Description = "filter logs by time range"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "time",
			Usage: "time to filter for",
		},
		cli.DurationFlag{
			Name:  "interval",
			Usage: "interval around the given time to filter",
			Value: time.Second,
		},
		cli.StringFlag{
			Name:  "time-format",
			Usage: "format of the given time. https://golang.org/pkg/time/#pkg-constants",
			Value: time.RFC3339,
		},
	}

	app.Action = func(c *cli.Context) error {
		return nil
	}
}
