package commons

import "github.com/sergiorb/iota-emulator/src/protocols"
import "github.com/sergiorb/iota-emulator/src/transports"
import "github.com/sergiorb/iota-emulator/src/generator"

type SHP struct {
	Schema	string	`json: schema`
	Host	string	`json: host`
	Port	uint16	`json: port`
}

type Http struct {
	Schema		string			`json: schema`
	Host		string			`json: host`
	Port		uint16			`json: port`
	Defaults	HttpDefaults	`json: defaults`
}

type HttpDefaults struct {
	Key			string	`json: key`
	Resource	string	`json: resource`
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
	Id			string					`json: id`
	Protocol	protocols.Protocol		`json: protocol`
	Transport	transports.Transport	`json: transport`
	Attributes	[]Attribute				`json: attributes`
}


type Attribute struct {
	Key				string					`json: key`
	Cron			string					`json: cron`
	ValueGenerator	generator.GeneratorType	`json: valueGenerator`
}
