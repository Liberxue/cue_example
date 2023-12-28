package config

import (
    "time"
)

Config : {
    server_name:        string | *"default-server"
    "server-timeout":   time.Duration | *"30s"
    endpoints:              [...string] & [...=~"^([0-9]{1,3}\\.){3}[0-9]{1,3}:[0-9]{1,5}$"]
    "help":      bool | *false
}

LiberConfig: [...Config]

liberConfig: LiberConfig & [{
    server_name:        "liber_test"
    endpoints: ["0.0.0.0:9090", "127.0.0.1:5672"]
    "help": false
}, {
    server_name:        "liber_test_server"
    endpoints: ["0.0.0.0:9091", "127.0.0.1:5672"]
    "help": false
}]
