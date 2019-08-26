package commons

import "github.com/sergiorb/iota-emulator/src/device"
import "github.com/sergiorb/iota-emulator/src/generator"

type SHP struct {
	Schema	string	`json: schema`
	Host	string	`json: host`
	Port	uint16	`json: port`
}

type Mqtt struct {
	Host	string	`json: host`
	Port	uint16	`json: port`
}

type Amqp struct {
	Host		string	`json: host`
	Port		uint16	`json: port`
	Exchange	string	`json: exchange`
	Queue		string	`json: queue`
}

type Device struct {
	Id			string				`json: id`
	Transport	device.Transport	`json: transport`
	Attributes	[]Attribute			`json: attributes`
}


type Attribute struct {
	Key				string					`json: key`
	Cron			string					`json: cron`
	ValueGenerator	generator.GeneratorType	`json: valueGenerator`
}
