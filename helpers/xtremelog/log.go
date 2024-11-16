package xtremelog

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/globalxtreme/gobaseconf/grpc/client"
	log2 "github.com/globalxtreme/gobaseconf/grpc/pkg/log"
	"github.com/globalxtreme/gobaseconf/helpers"
	"log"
	"os"
	"runtime/debug"
	"time"
)

func Info(content any) {
	logType := "INFO"
	if client.LogRPCActive {
		message, _ := json.Marshal(content)

		SendBugLog(&log2.LogRequest{
			Service: os.Getenv("SERVICE"),
			Type:    logType,
			Message: string(message),
		})
	} else {
		setOutput(logType, content)
	}
}

func Error(content any, bug bool) {
	debug.PrintStack()

	logType := "ERROR"
	if client.LogRPCActive {
		SendBugLog(&log2.LogRequest{
			Service: os.Getenv("SERVICE"),
			Type:    logType,
			Message: fmt.Sprintf("panic: %v", content),
			Stack:   debug.Stack(),
			Bug:     bug,
		})
	} else {
		setOutput("ERROR", fmt.Sprintf("panic: %v", content))
		setOutput("ERROR", string(debug.Stack()))
	}
}

func Debug(content any) {
	logType := "DEBUG"
	if client.LogRPCActive {
		message, _ := json.Marshal(content)

		SendBugLog(&log2.LogRequest{
			Service: os.Getenv("SERVICE"),
			Type:    logType,
			Message: string(message),
		})
	} else {
		setOutput("DEBUG", content)
	}
}

func SendBugLog(req *log2.LogRequest) (*log2.LGResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), client.LogRPCTimeout)
	defer cancel()

	return client.LogRPCClient.Log(ctx, req)
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
