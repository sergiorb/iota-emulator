package transports

type Transport string

const (
    HTTP Transport = "http"
    MQTT Transport = "mqtt"
    AMPQ Transport = "ampq"
    NONE Transport = "none"
)
