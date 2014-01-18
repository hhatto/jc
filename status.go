package main

import (
    "github.com/codegangsta/cli"
    "github.com/hhatto/nanairo"
    "bytes"
    "encoding/json"
    "fmt"
    "io"
)

type TaskInfo struct {
    Name string `json:"name"`
    Url string  `json:"url"`
}
type Queue struct {
    Task TaskInfo       `json:"task"`
    TaskId int          `json:"id"`
    EnQueueTime int     `json:"inQueueSince"`
    BuildStartTime int  `json:"buildableStartMilliseconds"`
}

type Executor struct {
    DisplayName string
    Progress int        // xx/100
}
type ComputerInfo struct {
    TotalExecutors int      `json:"totalExecutors"`
    BusyExecutors int       `json:"busyExecutors"`
    Executors []Executor    `json:"executors"`
}

func getExecutors(url string, dumpFlag bool) (ComputerInfo, error) {
    var r ComputerInfo
    client := NewClient(url)
    res, err := client.get("computer/api/json?depth=2")
    if err != nil {
        return r, err
    }
    defer res.Body.Close()

    if dumpFlag == true {
        buf := new(bytes.Buffer)
        buf.ReadFrom(res.Body)
        fmt.Println(buf.String())
        return r, nil
    }

    err = json.NewDecoder(io.Reader(res.Body)).Decode(&r)
    if err != nil {
        return r, err
    }

    return r, nil
}

func getQueue(url string, dumpFlag bool) ([]Queue, error) {
    client := NewClient(url)
    res, err := client.get("queue/api/json")
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if dumpFlag == true {
        buf := new(bytes.Buffer)
        buf.ReadFrom(res.Body)
        fmt.Println(buf.String())
        return nil, nil
    }

    var r struct {
        QueueItems []Queue `json:"items"`
    }
    err = json.NewDecoder(io.Reader(res.Body)).Decode(&r)
    if err != nil {
        return nil, err
    }

    return r.QueueItems, nil
}

func printJobQueue(url string, dumpFlag bool) {
    b := bytes.NewBufferString(" ")
    queueItems, _ := getQueue(url, dumpFlag)
    b.WriteString(fmt.Sprintf("build queue: %d", len(queueItems)))
    executors, _ := getExecutors(url, dumpFlag)

    b.WriteString(fmt.Sprintf(",  executor: %d/%d", executors.BusyExecutors, executors.TotalExecutors))
    if len(executors.Executors) > 0 {
        b = bytes.NewBufferString("\n ")
    }

    for _, queue := range queueItems {
        b.WriteString(nanairo.FgColor("#E0ffff", "✈"))
        b.WriteString("  - ")
        fmt.Println(queue)
        b.WriteString(nanairo.FgColor("#999999", queue.Task.Name))
    }

    for _, executor := range executors.Executors {
        b.WriteString(nanairo.FgColor("#ff6347", "✈"))
        b.WriteString("  -  ")
        b.WriteString(string(executor.Progress))
        b.WriteString("/100")
        b.WriteString(nanairo.FgColor("#999999", executor.DisplayName))
    }
    fmt.Println(b.String())
}

func printServerInfo(url string) {
    client := NewClient(url)
    res, err := client.head()
    if err != nil {
        fmt.Println("HEAD request error")
    }

    fmt.Println(" version:", res.Header["X-Jenkins"][0])
    fmt.Println(" server:", res.Header["Server"][0])
}

func status(c *cli.Context) {
    url := Config.Get(c.String("name"))
    fmt.Println(c.String("name"), "-", url)
    printJobQueue(url, c.Bool("dump"))
    fmt.Println("")
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
