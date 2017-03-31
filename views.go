package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hhatto/nanairo"
	"github.com/urfave/cli"
)

type View struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getViews(url string, dumpFlag bool) ([]View, error) {
	client := NewClient(url)
	res, err := client.get("api/json")
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
		ViewItems []View `json:"views"`
	}
	err = json.NewDecoder(io.Reader(res.Body)).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r.ViewItems, nil
}

func printView(url string, dumpFlag bool) {
	b := bytes.NewBufferString(" ")
	viewItems, _ := getViews(url, dumpFlag)

	for _, view := range viewItems {
		b.WriteString(fmt.Sprintf("\n  %s  %-20s - %s",
			nanairo.FgColor("#E0ffff", "👀"), view.Name, view.Url))
	}

	fmt.Println(b.String())
}

func viewsCommand(c *cli.Context) {
	url := Config.Get(c.String("name"))
	fmt.Println(c.String("name"), "-", url)
	printView(url, c.Bool("dump"))
	fmt.Println("")
}

var ViewsCommand = cli.Command{
	Name:   "views",
	Usage:  "print jenkins view list",
	Action: viewsCommand,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "default",
			Usage: "host key name(default is 'default')",
		},
	},
}
