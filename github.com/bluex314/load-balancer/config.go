package main

import (
	"sync"
)

const CONFIGFILE = "./config.json"

type Config struct {
	Proxy   Proxy     `json:"proxy"`
	Backend []Backend `json:"backends"`
}

// reverse proxy
type Proxy struct {
	Port string `json:port`
}

// servers which load balancer is transferred
type Backend struct {
	URL    string `json:"url"`
	IsDead bool
	mu     sync.RWMutex
}

var config Config

func readConfig() Config {
	readJsonFromFile(CONFIGFILE, &config)
	return config
}
