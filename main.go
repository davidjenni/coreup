package main

import (
	"os"

	"github.com/davidjenni/coreUp/cmds"
	"github.com/urfave/cli"
)

const appVersion = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Version = appVersion

	appOptions := cmds.GetAppOptions()
	app.Name = appOptions.Name
	app.Usage = appOptions.Usage
	app.Flags = appOptions.AppFlags
	app.Commands = appOptions.Commands
	app.CommandNotFound = appOptions.CommandNotFound

	app.Run(os.Args)
}
