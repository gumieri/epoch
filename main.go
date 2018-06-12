package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

func now(context *cli.Context) {
	epoch := time.Now().UnixNano()

	if context.Bool("milliseconds") {
		epoch = epoch / 1000000
	}

	fmt.Println(epoch)
}

func convert(context *cli.Context) {
	for _, arg := range context.Args() {
		epoch, err := strconv.ParseInt(arg, 10, 64)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(time.Unix(0, epoch))
	}
}

func main() {
	var err error

	app := cli.NewApp()
	app.Name = "epoch"
	app.Version = "0.0.1"
	app.ArgsUsage = "[timestamp]"
	app.Action = convert

	app.Commands = []cli.Command{
		{
			Name:   "now",
			Usage:  "Show the unix epoch timestamp from the moment that the command is executed",
			Action: now,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "milliseconds, m",
					Usage: "Timestamp as milliseconds",
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
