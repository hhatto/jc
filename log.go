package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jaytaylor/html2text"
	"github.com/urfave/cli"
)

type JobInfo struct {
	Number int    `json:"number"`
	Url    string `json:"url"`
	Class  string `json:"_class"`
}

func (j JobInfo) isPipelineJob() bool {
	if j.Class == "org.jenkinsci.plugins.workflow.job.WorkflowRun" {
		return true
	}
	return false
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

func printLogFromHTML(url string, jobName string, jobNumber int) error {
	client := NewClient(url)
	res, err := client.get(fmt.Sprintf("job/%s/%d/logText/progressiveHtml?start=0", jobName, jobNumber))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	newbuf := new(bytes.Buffer)

	textReader := bufio.NewReader(res.Body)
	for {
		b, isPrefix, err := textReader.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		buf := bytes.NewBuffer(b)
		if isPrefix {
			for {
				b, cont, err := textReader.ReadLine()
				if err != nil {
					if err != io.EOF {
						return err
					}
					break
				}
				buf.Write(b)
				if !cont {
					break
				}
			}
		}
		n := buf.String() + "<br />\n"
		newbuf.WriteString(n)
	}

	text, err := html2text.FromString(newbuf.String(), html2text.Options{PrettyTables: false})
	if err != nil {
		return err
	}
	fmt.Println(text)
	return nil
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
	if len(jobItems) == 0 {
		fmt.Println("not found job")
		return
	}

	job := jobItems[0]
	if job.isPipelineJob() {
		printLogFromHTML(url, jobName, job.Number)
	} else {
		printLog(url, jobName, job.Number)
	}
}

func logCommand(c *cli.Context) {
	url := Config.Get(c.String("name"))
	for _, jobName := range c.Args() {
		printLogWrapper(url, jobName, c.Bool("dump"))
	}
}

var LogCommand = cli.Command{
	Name:   "log",
	Usage:  "print job's log",
	Action: logCommand,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "default",
			Usage: "host key name(default is 'default')",
		},
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "verbose mode",
		},
		cli.BoolFlag{
			Name:  "dump, d",
			Usage: "dump raw json data",
		},
	},
}
