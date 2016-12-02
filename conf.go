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

	url := c.Args().First()
	if url == "" && c.Bool("rm") {
		fmt.Println("assign target name")
		return
	}

	if !c.Bool("set") && !c.Bool("dump") && !c.Bool("rm") {
		fmt.Println("please --set or --rm option")
		return
	}

	for i, _ := range Config.HostInfo {
		if Config.HostInfo[i].Name == c.String("name") {
			if c.Bool("rm") {
				Config.HostInfo = append(Config.HostInfo[:i+1], Config.HostInfo[i+1:]...)
			} else if c.Bool("set") {
				Config.HostInfo[i] = JcConfigHostInfo{Name: c.String("name"), Hostname: c.Args().First()}
				fmt.Println("add url:", c.Args().First())
			}
			Config.Save(defaultConfigFile)
			return
		}
	}

	Config.HostInfo = append(Config.HostInfo, JcConfigHostInfo{Name: c.String("name"), Hostname: c.Args().First()})
	Config.Save(defaultConfigFile)
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
