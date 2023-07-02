package console

import (
	"github.com/globalxtreme/gobaseconf/helpers"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"os"
	"strconv"
	"time"
)

type DeleteLogFileCommand struct{}

func (command DeleteLogFileCommand) Handle() {
	storageDir := os.Getenv("STORAGE_DIR") + "/logs/"

	logDays := 14
	logDaysEnv := os.Getenv("LOG_DAYS")
	if len(logDaysEnv) > 0 {
		logDays, _ = strconv.Atoi(logDaysEnv)
	}

	filename := time.Now().AddDate(0, 0, -logDays).Format(helpers.DateLayout()) + ".log"
	fullPath := storageDir + filename
	xtremelog.Debug(fullPath)

	_, err := os.Stat(fullPath)
	if err == nil {
		err := os.Remove(fullPath)
		if err != nil {
			xtremelog.Error(err)
		}
	}
}
