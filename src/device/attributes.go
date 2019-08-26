package device

type Attribute struct {
	Key				string	`json: key`
	ValueGenerator	string	`json: value_generator` // change to enum
	cron			string	`json: cron`
}
