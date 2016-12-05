package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func confCommand(c *cli.Context) {
	if c.Bool("dump") {
		dumpFormat := c.Args().First()
		if dumpFormat == "" {
			dumpFormat = "all"
		}
		Config.Dump(defaultConfigFile, dumpFormat)
		return
	}

	name := c.String("name")
	url := c.Args().First()

	if !c.Bool("set") && !c.Bool("dump") && !c.Bool("rm") {
		fmt.Println("please --set or --rm option")
		return
	}

	if _, ok := Config.HostInfo[name]; ok {
		if c.Bool("rm") {
			delete(Config.HostInfo, name)
			fmt.Println("delete name:", name, Config)
		} else if c.Bool("set") {
			Config.HostInfo[name] = JcConfigHostInfo{Name: name, Hostname: url}
			fmt.Println("rewrite url:", url)
		}
		Config.Save(defaultConfigFile)
	} else {
		if c.Bool("set") {
			Config.HostInfo[name] = JcConfigHostInfo{Name: name, Hostname: url}
			fmt.Println("add url:", url)
			Config.Save(defaultConfigFile)
		} else {
			fmt.Printf("name=%s is not found key\n", name)
		}
	}
}

var ConfCommand = cli.Command{
	Name:   "conf",
	Usage:  "config jc command setting param",
	Action: confCommand,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "default",
			Usage: "host key name(default is 'default')",
		},
		cli.BoolFlag{
			Name:  "set, s",
			Usage: "add (or overwrite) host key name",
		},
		cli.BoolFlag{
			Name:  "rm",
			Usage: "remove host key name",
		},
		cli.BoolFlag{
			Name:  "dump, d",
			Usage: "print configuration (all, list). default is 'all'",
		},
	},
}
