package main

import (
    "github.com/codegangsta/cli"
    "fmt"
)


func printServerInfo(hostname string) {
    client := NewClient(hostname)
    res, err := client.head()
    if err != nil {
        fmt.Println("HEAD request error")
    }

    fmt.Println("version:", res.Header["X-Jenkins"][0])
    fmt.Println("server:", res.Header["Server"][0])
}

func status(c *cli.Context) {
    if len(c.Args()) == 0 {
        fmt.Println("set target hostname")
        return
    }

    printServerInfo(c.Args()[0])
}

var Status = cli.Command {
    Name: "status",
    Usage: "print jenkins host status",
    Action: status,
}
