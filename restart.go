package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli"
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

func restartCommand(c *cli.Context) {
	url := Config.Get(c.String("name"))
	executeRestart(url, c.Bool("force"))
}

var RestartCommand = cli.Command{
	Name:   "restart",
	Usage:  "restart jenkins server",
	Action: restartCommand,
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
			Name:  "force, f",
			Usage: "force restart (usually SafeRestat)",
		},
	},
}
