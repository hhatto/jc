package main

import (
    "github.com/codegangsta/cli"
    "os"
)

const VERSION = "0.0.1"


func main() {
    app := cli.NewApp()
    app.Name = "jc"
    app.Usage = "j + c = jenkins"
    app.Version = VERSION
    app.Commands = subCommands

    err := app.Run(os.Args)
    if err != nil {
        os.Exit(1)
    }
    os.Exit(0)
}
