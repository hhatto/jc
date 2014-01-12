package main

import (
    "github.com/codegangsta/cli"
    "fmt"
)


func printServerInfo(url string) {
    client := NewClient(url)
    res, err := client.head()
    if err != nil {
        fmt.Println("HEAD request error")
    }

    fmt.Println("version:", res.Header["X-Jenkins"][0])
    fmt.Println("server:", res.Header["Server"][0])
}

func status(c *cli.Context) {
    url := Config.Get(c.String("name"))
    fmt.Println(c.String("name"), "-", url)
    printServerInfo(url)
}

var Status = cli.Command {
    Name: "status",
    Usage: "print jenkins host status",
    Action: status,
    Flags: []cli.Flag {
        cli.StringFlag {
            "name, n", "default",
            "host key name(default is 'default')",},
    },
}
