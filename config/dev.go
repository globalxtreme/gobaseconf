package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

var DevMode bool

func InitDevMode(force ...bool) {
	isDev := DevMode
	if len(force) > 0 && force[0] {
		isDev = true
	}

	if isDev {
		fmt.Println("Running in development mode..")
		err := godotenv.Load()
		if err != nil {
			panic(err.Error())
		}
	}
}
