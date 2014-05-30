package main

import (
    "bytes"
    "fmt"
    "github.com/codegangsta/cli"
)

func executeJob(url string, jobName string) error {
    client := NewClient(url)
    res, err := client.get(fmt.Sprintf("job/%s/build", jobName))
    if err != nil {
        return err
    }
    defer res.Body.Close()

    if res.StatusCode == 200 || res.StatusCode == 201 {
        fmt.Println(fmt.Sprintf("OK. building '%s'", jobName))
    } else {
        buf := new(bytes.Buffer)
        buf.ReadFrom(res.Body)
        fmt.Println(buf.String())
    }
    return nil
}

func buildWrapper(url string, jobName string, dumpFlag bool) {
    executeJob(url, jobName)
}

func build(c *cli.Context) {
    url := Config.Get(c.String("name"))
    for _, jobName := range c.Args() {
        buildWrapper(url, jobName, c.Bool("dump"))
    }
}

var BuildCommand = cli.Command{
    Name:   "build",
    Usage:  "build job",
    Action: build,
    Flags: []cli.Flag{
        cli.StringFlag{
            "name, n", "default",
            "host key name(default is 'default')"},
        cli.BoolFlag{
            "verbose, v", "verbose mode",
        },
        cli.BoolFlag{
            "dump, d", "dump raw json data",
        },
    },
}
