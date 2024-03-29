package helpers

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func FullDateTimeLayout() string {
	return "02/01/2006 15:04:05"
}

func DateLayout() string {
	return "02/01/2006"
}

func TimeLayout() string {
	return "15:04:05"
}

func RandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randomBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		randomBytes[i] = chars[rand.Intn(len(chars))]
	}

	return string(randomBytes) + strconv.FormatInt(time.Now().UnixNano(), 10)
}

func CheckAndCreateDirectory(path string) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
		}
	}
}

func SetStorageDir(path ...string) string {
	storagePath := os.Getenv("STORAGE_DIR")
	if len(storagePath) == 0 {
		storagePath = "storages"
	}

	if len(path) > 0 {
		storagePath += "/" + path[0]
	}

	return storagePath
}

func SetStorageAppDir(path ...string) string {
	appDir := "app"
	if len(path) > 0 {
		appDir += "/" + path[0]
	}

	return SetStorageDir(appDir)
}

func SetStorageAppPublicDir(path ...string) string {
	publicDir := "app/public"
	if len(path) > 0 {
		publicDir += "/" + path[0]
	}

	return SetStorageDir(publicDir)
}

func StringToArrayInt(text string) []int {
	var array []int
	texts := strings.Split(text, ",")
	for _, value := range texts {
		item, _ := strconv.Atoi(value)
		array = append(array, item)
	}

	return array
}

func StringToArrayString(text string) []string {
	var array []string
	texts := strings.Split(text, ",")
	for _, value := range texts {
		array = append(array, value)
	}

	return array
}
