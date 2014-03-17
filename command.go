package main

import (
    "github.com/codegangsta/cli"
    "net/http"
)


type Client struct {
    baseUrl string
    *http.Client
}

var subCommands = []cli.Command {
    JobsCommand, JobCommand, StatusCommand, ConfCommand,
}


func NewClient(url string) *Client {
    return &Client{url + "/", http.DefaultClient}
}

func (c *Client) get(path string) (*http.Response, error) {
    req, err := http.NewRequest("GET", c.baseUrl + path, nil)
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
