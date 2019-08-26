package v1

import "github.com/sergiorb/iota-emulator/src/config/commons"

type IotaJsonV1 struct {
	Http	commons.SHP		`json: http`
	Mqtt	commons.Mqtt	`json: mqtt`
	Amqp	commons.Amqp	`json: amqp`
}
