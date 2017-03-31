package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hhatto/nanairo"
	"github.com/urfave/cli"
)

type TaskInfo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type Queue struct {
	Task           TaskInfo `json:"task"`
	TaskId         int      `json:"id"`
	EnQueueTime    int      `json:"inQueueSince"`
	BuildStartTime int      `json:"buildableStartMilliseconds"`
}

type ExecutorDetail struct {
	DisplayName string `json:"fullDisplayName"`
}
type Executor struct {
	Detail   ExecutorDetail `json:"currentExecutable"`
	Progress int            `json:"progress"`
}
type ComputerInfo struct {
	Executors       []Executor `json:"executors"`
	OneOffExecutors []Executor `json:"oneOffExecutors"`
}
type AllExecutor struct {
	TotalExecutors int            `json:"totalExecutors"`
	BusyExecutors  int            `json:"busyExecutors"`
	ComputerInfos  []ComputerInfo `json:"computer"`
}

func getExecutors(url string, dumpFlag bool) (AllExecutor, error) {
	var r AllExecutor
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

	for _, queue := range queueItems {
		b.WriteString(fmt.Sprintf("\n  %s  - %-20s %s",
			nanairo.FgColor("#E0ffff", "✈ ⇥"), queue.Task.Name,
			nanairo.FgColor("#666666", "(in build queue)")))
	}

	if len(executors.ComputerInfos) <= 0 {
		fmt.Println("jenkins invalid state")
		return
	}

	for _, info := range executors.ComputerInfos {
		for _, executor := range info.Executors {
			if executor.Progress < 0 {
				continue
			}
			if executor.Detail.DisplayName == "" {
				continue
			}
			b.WriteString(fmt.Sprintf("\n  %s  - %-20s %s",
				nanairo.FgColor("#ff6347", "✈ ➟"), executor.Detail.DisplayName,
				nanairo.FgColor("#666666", fmt.Sprintf("(%d/100)", executor.Progress))))
		}
		for _, executor := range info.OneOffExecutors {
			if executor.Progress < 0 {
				continue
			}
			if executor.Detail.DisplayName == "" {
				continue
			}
			b.WriteString(fmt.Sprintf("\n  %s  - %-20s %s",
				nanairo.FgColor("#ff6347", "✈ ➟"), executor.Detail.DisplayName,
				nanairo.FgColor("#666666", fmt.Sprintf("(%d/100)", executor.Progress))))
		}
	}
	fmt.Println(b.String())
}

func printServerInfo(url string) {
	client := NewClient(url)
	res, err := client.head()
	if err != nil {
		fmt.Println("HEAD request error")
		return
	}

	fmt.Println(" version:", res.Header["X-Jenkins"][0])
	fmt.Println(" server:", res.Header["Server"][0])
}

func statusCommand(c *cli.Context) {
	url := Config.Get(c.String("name"))
	fmt.Println(c.String("name"), "-", url)
	printJobQueue(url, c.Bool("dump"))
	fmt.Println("")
	printServerInfo(url)
}

var StatusCommand = cli.Command{
	Name:   "status",
	Usage:  "print jenkins host status",
	Action: statusCommand,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "default",
			Usage: "host key name(default is 'default')",
		},
	},
}
