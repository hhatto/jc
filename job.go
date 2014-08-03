package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/hhatto/nanairo"
	"io"
	"strconv"
	"time"
)

type JobBuildInfo struct {
	Number    int    `json:"number"`
	Result    string `json:"result"`
	Timestamp int64  `json:"timestamp"`
	Duration  int    `json:"duration"`
}

func getJobInfo(url string, jobName string, dumpFlag bool) ([]JobBuildInfo, error) {
	client := NewClient(url)
	res, err := client.get(fmt.Sprintf("job/%s/api/json?depth=1", jobName))
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
		JobInfo []JobBuildInfo `json:"builds"`
	}
	err = json.NewDecoder(io.Reader(res.Body)).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r.JobInfo, nil
}

func printJobDetail(url string, jobName string, dumpFlag bool) {
	b := bytes.NewBufferString("")
	jobItems, _ := getJobInfo(url, jobName, dumpFlag)

	if len(jobItems) > 0 {
		b.WriteString(fmt.Sprintf("[%s]\n", jobName))
	} else {
		b.WriteString(fmt.Sprintf("'%s' job is not found\n", jobName))
	}

	for cnt, job := range jobItems {
		if cnt >= 5 {
			break
		}
		resultNumber := nanairo.FgColor("#0c0", strconv.Itoa(job.Number))
		if job.Result != "SUCCESS" {
			resultNumber = nanairo.FgColor("#c00", strconv.Itoa(job.Number))
		}
		humanReadableDurationValue := fmt.Sprintf("%ds", job.Duration/1000)
		if job.Duration/1000 >= 60 {
			humanReadableDurationValue = fmt.Sprintf("%dm%ds", job.Duration/1000/60, job.Duration/1000%60)
		}
		b.WriteString(fmt.Sprintf("  [%4s] %s (%6s)\n", resultNumber,
			time.Unix(int64(job.Timestamp)/1000, 0),
			humanReadableDurationValue))
	}

	fmt.Print(b.String())
}

func job(c *cli.Context) {
	url := Config.Get(c.String("name"))
	for _, jobName := range c.Args() {
		printJobDetail(url, jobName, c.Bool("dump"))
	}
}

var JobCommand = cli.Command{
	Name:   "job",
	Usage:  "print job detail",
	Action: job,
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
