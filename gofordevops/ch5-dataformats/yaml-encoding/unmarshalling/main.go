package main

import (
	"fmt"
	"time"

	"github.com/go-yaml/yaml"
)

var yamlFile = []byte(`
jobs:
- name: Clear tmp
  interval: 24h0m0s
  whatever: is not in the job type
  cmd: rm -rf /tmp
`)

type Config struct {
	Jobs []Job
}

type Job struct {
	Name     string
	Interval time.Duration
	Cmd      string
}

func main() {

	c := Config{}

	if err := yaml.UnmarshalStrict(yamlFile, &c); err != nil {
		panic(err)
	}

	// We check that this key is set so tat our type assertion below will not panic.
	for _, j := range c.Jobs {
		fmt.Printf("Name: %s\nInterval: %s", j.Name, j.Interval)
	}
}
