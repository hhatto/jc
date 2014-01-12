package main

import (
    "github.com/codegangsta/cli"
    "os"
)

const VERSION = "0.0.1"


func setup() {
    _, err := os.Stat(defaultConfigFile)
    if os.IsNotExist(err) {
        hostInfo := make([]JcConfigHostInfo, 1)
        Config.HostInfo = hostInfo
        Config.HostInfo[0] = JcConfigHostInfo{Name: "default", Hostname: ""}
        Config.Save(defaultConfigFile)
    }
    Config.Load(defaultConfigFile)
}

func main() {
    app := cli.NewApp()
    app.Name = "jc"
    app.Usage = "j + c = jenkins"
    app.Version = VERSION
    app.Commands = subCommands

    setup()

    err := app.Run(os.Args)
    if err != nil {
        os.Exit(1)
    }
    os.Exit(0)
}
