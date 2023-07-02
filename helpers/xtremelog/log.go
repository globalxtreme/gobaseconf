package xtremelog

import (
	"fmt"
	"github.com/globalxtreme/gobaseconf/helpers"
	"log"
	"os"
	"time"
)

func Info(content any) {
	setOutput("INFO", content)
}

func Error(content any) {
	setOutput("ERROR", content)
}

func Debug(content any) {
	setOutput("DEBUG", content)
}

func setOutput(action string, error any) {
	storageDir := os.Getenv("STORAGE_DIR") + "/logs"
	helpers.CheckAndCreateDirectory(storageDir)

	filename := time.Now().Format(helpers.DateLayout()) + ".log"
	file, err := os.OpenFile(storageDir+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(fmt.Sprintf("[%s]:", action), error)
}
