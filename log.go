package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io"
)

type JobInfo struct {
	Number int    `json:"number"`
	Url    string `json:"url"`
}

func getJobInfoSimple(url string, jobName string, dumpFlag bool) ([]JobInfo, error) {
	client := NewClient(url)
	res, err := client.get(fmt.Sprintf("job/%s/api/json", jobName))
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
		JobInfo []JobInfo `json:"builds"`
	}
	err = json.NewDecoder(io.Reader(res.Body)).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r.JobInfo, nil
}

func printLog(url string, jobName string, jobNumber int) error {
	client := NewClient(url)
	res, err := client.get(fmt.Sprintf("job/%s/%d/logText/progressiveText?start=0", jobName, jobNumber))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	fmt.Println(buf.String())
	return nil
}

func printLogWrapper(url string, jobName string, dumpFlag bool) {
	b := bytes.NewBufferString("")
	jobItems, _ := getJobInfoSimple(url, jobName, dumpFlag)

	b.WriteString(fmt.Sprintf("[%s]\n", jobName))
	if len(jobItems) != 0 {
		printLog(url, jobName, jobItems[0].Number)
	}
}

func log(c *cli.Context) {
	url := Config.Get(c.String("name"))
	for _, jobName := range c.Args() {
		printLogWrapper(url, jobName, c.Bool("dump"))
	}
}

var LogCommand = cli.Command{
	Name:   "log",
	Usage:  "print job's log",
	Action: log,
	Flags: []cli.Flag{
		cli.StringFlag{
			"name, n",
			"default",
			"host key name(default is 'default')",
			"",
		},
		cli.BoolFlag{
			"verbose, v",
			"verbose mode",
			"",
		},
		cli.BoolFlag{
			"dump, d",
			"dump raw json data",
			"",
		},
	},
}
