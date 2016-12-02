package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"

	"github.com/hhatto/nanairo"
	"github.com/urfave/cli"
)

type healthReport struct {
	Score int `json:"score"`
}

type Job struct {
	Name         string         `json:"name"`
	Url          string         `json:"url"`
	Color        string         `json:"color"`
	HealthReport []healthReport `json:"healthReport"`
}

func getJobs(targetUrl string, dumpFlag bool) ([]Job, error) {
	client := NewClient(targetUrl)
	res, err := client.get("api/json?depth=1")
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
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
		Jobs []Job `json:"jobs"`
	}
	err = json.NewDecoder(io.Reader(res.Body)).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r.Jobs, nil
}

func jobsCommand(c *cli.Context) {
	targetHost := Config.Get(c.String("name"))
	u, err := url.Parse(targetHost)
	if err != nil {
		log.Fatal(err)
	}
	isViewAccess := c.Args().Present()

	if isViewAccess {
		u.Path = fmt.Sprintf("view/%s", c.Args().First())
	}
	jobs, _ := getJobs(u.String(), c.Bool("dump"))
	if c.Bool("dump") == true || len(jobs) == 0 {
		if isViewAccess {
			fmt.Println(fmt.Sprintf("'%s' view is not exists", c.Args().First()))
		}
		return
	}
	fmt.Println(c.String("name"), "-", targetHost)
	if isViewAccess {
		fmt.Println(fmt.Sprintf("[view:%s]", c.Args().First()))
	}
	for _, job := range jobs {
		// S
		var j = bytes.NewBufferString("  ")
		if job.Color == "blue" {
			j.WriteString(nanairo.FgColor("#0c0", "✔"))
		} else if job.Color == "disabled" {
			j.WriteString(nanairo.FgColor("#666", "✔"))
		} else if job.Color == "blue_anime" {
			j.WriteString(nanairo.FgColor("#0cc", "➟"))
		} else {
			j.WriteString(nanairo.FgColor("#c00", "✔"))
		}
		j.WriteString("  ")

		// W
		if len(job.HealthReport) == 0 {
			j.WriteString(nanairo.FgColor("#aaa", "⁇"))
		} else if job.HealthReport[0].Score >= 80 {
			j.WriteString(nanairo.FgColor("#f90", "☀"))
		} else if job.HealthReport[0].Score >= 20 {
			j.WriteString(nanairo.FgColor("#999", "☁"))
		} else {
			j.WriteString(nanairo.FgColor("#03c", "☂"))
		}

		// Name
		j.WriteString(fmt.Sprintf("  %s", job.Name))

		fmt.Println(j.String())
	}
}

var JobsCommand = cli.Command{
	Name:   "jobs",
	Usage:  "print status for all jobs",
	Action: jobsCommand,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "default",
			Usage: "host key name(default is 'default')",
		},
		cli.BoolFlag{
			Name:  "dump, d",
			Usage: "dump raw json data",
		},
	},
}
