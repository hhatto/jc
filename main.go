package main

import (
	"os"

	"github.com/urfave/cli"
)

const VERSION = "0.1"

func setup() {
	_, err := os.Stat(defaultConfigFile)
	if os.IsNotExist(err) {
		Config.HostInfo = make(map[string]JcConfigHostInfo)
		Config.HostInfo["default"] = JcConfigHostInfo{Name: "default", Hostname: ""}
		Config.Save(defaultConfigFile)
	}
	Config.Load(defaultConfigFile)
}

func main() {
	app := cli.NewApp()
	app.Name = "jc"
	app.Usage = "j + c = jenkins cli"
	app.Version = VERSION
	app.Commands = subCommands

	setup()

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
