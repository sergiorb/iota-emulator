package iotajson

import agent "github.com/sergiorb/iota-emulator/src/config/iota-json/v1"
import "github.com/sergiorb/iota-emulator/src/config/commons"

type Config struct {
	Agent	agent.IotaJsonV1	`json: agent`
	Devices []commons.Device	`json: devices`
}
