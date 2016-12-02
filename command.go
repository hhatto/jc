package main

import (
	"net/http"

	"github.com/urfave/cli"
)

type Client struct {
	baseUrl string
	*http.Client
}

var subCommands = []cli.Command{
	BuildCommand, JobsCommand, JobCommand, StatusCommand, ConfCommand, LogCommand,
	RestartCommand,
}

func NewClient(url string) *Client {
	return &Client{url + "/", http.DefaultClient}
}

func (c *Client) get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.baseUrl+path, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) head() (*http.Response, error) {
	req, err := http.NewRequest("HEAD", c.baseUrl, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) post(path string) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.baseUrl+path, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
