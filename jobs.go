package main

import (
    "github.com/codegangsta/cli"
    "github.com/hhatto/nanairo"
    "fmt"
    "io"
    "encoding/json"
)

type Job struct {
    Name string `json:"name"`
    Url string `json:"url"`
    Color string `json:"color"`
}

func getJobs(hostname string) ([]Job, error) {
    client := NewClient(hostname)
    res, err := client.get("api/json")
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    var r struct {
        Jobs []Job `json:"jobs"`
    }
    err = json.NewDecoder(io.Reader(res.Body)).Decode(&r)
    if err != nil {
        return nil, err
    }

    return r.Jobs, nil
}

func jobs(c *cli.Context) {
    if len(c.Args()) > 0 {
        jobs, _ := getJobs(c.Args()[0])
        for _, job := range jobs {
            if job.Color == "blue" {
                fmt.Println(nanairo.FgColor("#0c0", "✔") + " " + job.Name)
            } else {
                fmt.Println(nanairo.FgColor("#c00", "✔") + " " + job.Name)
            }
        }
    } else {
        fmt.Println("set target hostname")
    }
}

var Jobs = cli.Command {
    Name: "jobs",
    Usage: "print status for all jobs",
    Action: jobs,
}
