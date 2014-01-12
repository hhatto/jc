package main

import (
    "github.com/codegangsta/cli"
    "fmt"
)


func conf(c *cli.Context) {
    if c.Bool("dump") {
        Config.Dump(defaultConfigFile)
        return
    }

    if len(c.Args()) == 0 {
        fmt.Println("assign target url")
        return
    }

    for i, _ := range Config.HostInfo {
        if Config.HostInfo[i].Name == c.String("name") {
            Config.HostInfo[i] = JcConfigHostInfo{Name: c.String("name"), Hostname: c.Args().First()}
            Config.Save(defaultConfigFile)
            return
        }
    }

    Config.HostInfo = append(Config.HostInfo, JcConfigHostInfo{Name: c.String("name"), Hostname: c.Args().First()})
    Config.Save(defaultConfigFile)
}

var Conf = cli.Command {
    Name: "conf",
    Usage: "config jc command setting param",
    Action: conf,
    Flags: []cli.Flag {
        cli.StringFlag {
            "name, n", "default",
            "host key name(default is 'default')",},
        cli.BoolFlag {
            "dump, d",
            "print configuration",},
    },
}
