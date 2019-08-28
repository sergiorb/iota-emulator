package jobs

import "fmt"
import "time"
import "encoding/json"
import "github.com/sergiorb/iota-emulator/src/generator"
import "github.com/sergiorb/iota-emulator/src/logs"
import "github.com/sergiorb/iota-emulator/src/transports"
import "github.com/sergiorb/iota-emulator/src/protocols"
import "github.com/sergiorb/iota-emulator/src/agents"

type AttributeJob struct {

	DeviceId		string
	Key				string
	Cron			string
	
	ProtocolType	protocols.Protocol
	TransportType	transports.Transport

	Generator 		generator.Generator

	Agent			agents.Agent
}

func (j *AttributeJob) Run() {

	var value = j.Generator.Generate();

	js, _ := json.Marshal(logs.AttributeJobLog{j.DeviceId, j.Key, value, time.Now()});

	fmt.Printf("%v\n", string(js));

	request := agents.Request{
		DeviceId:	j.DeviceId,
		Attribute:	j.Key,
		Value:		value,
	}
	
	j.Agent.Send(request, j.TransportType);
}
