package device

type Transport string

const (
    HTTP Transport = "HTTP"
    MQTT Transport = "MQTT"
	AMPQ Transport = "AMPQ"
)
