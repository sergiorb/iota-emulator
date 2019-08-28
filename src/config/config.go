package config

import "github.com/sergiorb/iota-emulator/src/config/commons"
import jsonAgent "github.com/sergiorb/iota-emulator/src/config/iota-json/v1"

type Config struct {
	Agents	Agents				`json: agents`
	Devices []commons.Device	`json: devices`
}

type Agents struct {
	Json	jsonAgent.IotaJsonV1	`json: json`
}
