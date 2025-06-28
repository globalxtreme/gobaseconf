package config

import (
	"os"
	"strconv"
)

var (
	HostFull    string
	Host        string
	Port        string
	Protocol    string
	ServiceName string
)

func SetHost() {
	Protocol = "http"
	SSL, _ := strconv.ParseBool(os.Getenv("USE_SSL"))
	if SSL == true {
		Protocol = "https"
	}

	Host = os.Getenv("DOMAIN")
	Port = os.Getenv("PORT")

	HostFull = Protocol + "://" + Host
	if SSL == false {
		HostFull += ":" + Port
	}
}

func GetHostFull() string {
	if len(HostFull) == 0 {
		SetHost()
	}

	return HostFull
}

func GetServiceName() string {
	if ServiceName == "" {
		ServiceName = os.Getenv("SERVICE")
	}

	return ServiceName
}
