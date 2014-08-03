package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func executeRestart(url string, forceFlag bool) error {
	var path string
	var mode string
	if forceFlag {
		path = "restart"
		mode = "force"
	} else {
		path = "safeRestart"
		mode = "safe"
	}

	fmt.Printf("url:%s\n(%s)restart ok? [y/n]: ", url, mode)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Text()[0] != 'y' {
		return nil
	}

	client := NewClient(url)
	res, err := client.post(path)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fmt.Printf("%s ok\n", path)

	return nil
}

func restart(c *cli.Context) {
	url := Config.Get(c.String("name"))
	executeRestart(url, c.Bool("force"))
}

var RestartCommand = cli.Command{
	Name:   "restart",
	Usage:  "restart jenkins server",
	Action: restart,
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
			"force, f",
			"force restart (usually SafeRestat)",
			"",
		},
	},
}
