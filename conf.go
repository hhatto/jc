package main

import (
    "github.com/codegangsta/cli"
    "fmt"
)


func conf(c *cli.Context) {
    if c.String("dump") != "" {
        Config.Dump(defaultConfigFile, c.String("dump"))
        return
    }

    url := c.Args().First()
    if url == "" && c.String("rm") != "" {
        fmt.Println("assign target url")
        return
    }

    for i, _ := range Config.HostInfo {
        if Config.HostInfo[i].Name == c.String("name") {
            if c.String("rm") != "" {
                // FIXME
                Config.HostInfo = append(Config.HostInfo[:i+1], Config.HostInfo[i+1:]...)
                fmt.Println(len(Config.HostInfo))
            } else {
                Config.HostInfo[i] = JcConfigHostInfo{Name: c.String("name"), Hostname: c.Args().First()}
                fmt.Println("url:", c.Args().First())
            }
            Config.Save(defaultConfigFile)
            return
        }
    }

    Config.HostInfo = append(Config.HostInfo, JcConfigHostInfo{Name: c.String("name"), Hostname: c.Args().First()})
    Config.Save(defaultConfigFile)
}

var ConfCommand = cli.Command {
    Name: "conf",
    Usage: "config jc command setting param",
    Action: conf,
    Flags: []cli.Flag {
        cli.StringFlag {
            "name, n", "default",
            "host key name(default is 'default')",},
        cli.StringFlag {
            "rm", "",
            "remove host key name",},
        cli.StringFlag {
            "dump, d", "",
            "print configuration (all, list)",},
    },
}
