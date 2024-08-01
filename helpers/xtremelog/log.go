package xtremelog

import (
	"encoding/json"
	"fmt"
	"github.com/globalxtreme/gobaseconf/grpc/client"
	"github.com/globalxtreme/gobaseconf/grpc/pkg/bug"
	"github.com/globalxtreme/gobaseconf/helpers"
	"log"
	"os"
	"runtime/debug"
	"time"
)

func Info(content any) {
	logType := "INFO"
	if client.BugRPCActive {
		message, _ := json.Marshal(content)

		client.BugLog(&bug.LogRequest{
			Service: os.Getenv("SERVICE"),
			Type:    logType,
			Message: message,
		})
	} else {
		setOutput(logType, content)
	}
}

func Error(content any) {
	debug.PrintStack()

	logType := "ERROR"
	if client.BugRPCActive {
		client.BugLog(&bug.LogRequest{
			Service: os.Getenv("SERVICE"),
			Type:    logType,
			Title:   fmt.Sprintf("panic: %v", content),
			Message: debug.Stack(),
		})
	} else {
		setOutput("ERROR", fmt.Sprintf("panic: %v", content))
		setOutput("ERROR", string(debug.Stack()))
	}
}

func Debug(content any) {
	logType := "DEBUG"
	if client.BugRPCActive {
		message, _ := json.Marshal(content)

		client.BugLog(&bug.LogRequest{
			Service: os.Getenv("SERVICE"),
			Type:    logType,
			Message: message,
		})
	} else {
		setOutput("DEBUG", content)
	}
}

func setOutput(action string, error any) {
	storageDir := os.Getenv("STORAGE_DIR") + "/logs"
	helpers.CheckAndCreateDirectory(storageDir)

	filename := time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile(storageDir+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(fmt.Sprintf("[%s]:", action), error)
}
