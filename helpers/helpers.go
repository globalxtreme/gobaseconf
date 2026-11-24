package helpers

import (
	"google.golang.org/grpc/status"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Numeric interface {
	uint | int | int8 | int16 | int32 | int64 | float32 | float64
}

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
	re := regexp.MustCompile(`[^0-9\-.]`)
	text = re.ReplaceAllString(text, " ")

	re = regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")

	words := strings.Split(text, " ")

	var result []int
	for _, word := range words {
		if word != "" {
			number, _ := strconv.Atoi(word)
			result = append(result, number)
		}
	}

	return result
}

func StringToArrayNumeric[N Numeric](text string) []N {
	re := regexp.MustCompile(`[^0-9\-.]`)
	text = re.ReplaceAllString(text, " ")

	re = regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")

	words := strings.Split(text, " ")

	var result []N
	for _, word := range words {
		if word != "" {
			number, _ := strconv.ParseFloat(word, 64)
			result = append(result, N(number))
		}
	}

	return result
}

func StringToArrayString(text string) []string {
	re := regexp.MustCompile(`[^A-Za-z0-9\-.]`)
	text = re.ReplaceAllString(text, " ")

	re = regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")

	words := strings.Split(text, " ")

	var result []string
	for _, word := range words {
		if word != "" {
			result = append(result, word)
		}
	}

	return result
}

func ToInt(text string) int {
	value, _ := strconv.Atoi(text)
	return value
}

func ToBool(text string) bool {
	value, _ := strconv.ParseBool(text)
	return value
}

func ToFloat64(text string) float64 {
	value, _ := strconv.ParseFloat(text, 64)
	return value
}

func ToNumeric[N Numeric](text string) N {
	number, _ := strconv.ParseFloat(text, 64)
	return N(number)
}

func GRPCErrorMessage(err error) string {
	st, ok := status.FromError(err)
	if ok {
		return st.Message()
	} else {
		return err.Error()
	}
}
