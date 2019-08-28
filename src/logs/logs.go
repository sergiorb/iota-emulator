package logs

import "time"

type AttributeJobLog struct {
	DeviceId	string		`json: device_id`
	Key			string		`json: key`
	Value		string		`json: value`
	Date		time.Time	`json: date`
}
