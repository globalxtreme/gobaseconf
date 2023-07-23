package helpers

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func FullDateTimeLayout() string {
	return "2006-01-02 15:04:05"
}

func DateLayout() string {
	return "2006-01-02"
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