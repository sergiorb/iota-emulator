package config

import jsonAgent "github.com/sergiorb/iota-emulator/src/config/iota-json"

type Config struct {
	Json	jsonAgent.Config	`json: json`
}
