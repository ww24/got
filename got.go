package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/ww24/got/cmd"
)

var app = cli.NewApp()

func init() {
	app.Name = "got"
	app.Version = "0.1.0"
	app.Usage = "Go Twitter utility tools"
	app.Action = action
	app.Commands = []cli.Command{
		cmd.FetchImages,
	}
}

func main() {
	app.Run(os.Args)
}

func action(ctx *cli.Context) {
	cli.ShowAppHelp(ctx)
	os.Exit(1)
}
