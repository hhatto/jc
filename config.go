package main

import (
	"encoding/json"
	"fmt"
	"io"
	llog "log"
	"os"
	"path/filepath"
)

var (
	defaultConfigFile = filepath.Join(os.Getenv("HOME"), ".config", "jc")
	Config            = NewJcConfig()
)

type JcConfigHostInfo struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
}

type JcConfig struct {
	HostInfo []JcConfigHostInfo `json:"hosts"`
}

func NewJcConfig() *JcConfig {
	c := new(JcConfig)
	return c
}

func (conf *JcConfig) Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	for {
		if err := dec.Decode(conf); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}

func (conf *JcConfig) Dump(filename string, format string) error {
	if format == "list" {
		for _, e := range conf.HostInfo {
			fmt.Println(fmt.Sprintf("%s,%s", e.Name, e.Hostname))
		}
		return nil
	}

	enc, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(enc))
	return err
}

func (conf *JcConfig) Save(filename string) error {
	err := os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	enc, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(enc)
	return err
}

func (conf *JcConfig) Get(key string) string {
	for _, info := range conf.HostInfo {
		if info.Name == key {
			return info.Hostname
		}
	}

	llog.Fatalf("not found key: %s", key)
	return ""
}
