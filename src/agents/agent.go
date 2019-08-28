package agents

import "github.com/sergiorb/iota-emulator/src/transports"

type Request struct {
	Key			string
	Resource	string
	DeviceId	string
	Attribute	string
	Value		interface{}
}

type Agent interface {

	Send(req Request, transport transports.Transport);
}
