package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

var DevMode bool

func InitDevMode() {
	if DevMode {
		fmt.Println("Running in development mode..")
		err := godotenv.Load()
		if err != nil {
			panic(err.Error())
		}
	}
}
