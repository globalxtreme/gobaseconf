package command

import (
	"fmt"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
	"time"
)

type DeleteLogFileCommand struct {
	FromSchedule bool
}

func (c *DeleteLogFileCommand) Command(cmd *cobra.Command) {
	cmd.AddCommand(&cobra.Command{
		Use:  "delete-log-file",
		Long: "Delete log file command",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode(c.FromSchedule)

			c.Handle()
		},
	})
}

func (c *DeleteLogFileCommand) Handle() {
	storageDir := os.Getenv("STORAGE_DIR") + "/logs/"

	logDays := 14
	logDaysEnv := os.Getenv("LOG_DAYS")
	if len(logDaysEnv) > 0 {
		logDays, _ = strconv.Atoi(logDaysEnv)
	}
	logDays++

	now := time.Now()
	dateLimitRemove, _ := time.ParseInLocation("2006-01-02", now.Format("2006-01-02"), now.Location())
	dateLimitRemove = dateLimitRemove.AddDate(0, 0, -logDays)

	logDirs, _ := os.ReadDir(storageDir)
	for _, logDir := range logDirs {
		if logDir.IsDir() {
			continue
		}

		logDate, err := time.Parse("2006-01-02", strings.Replace(logDir.Name(), ".log", "", 1))
		if err != nil {
			continue
		}

		if logDate.After(dateLimitRemove) {
			continue
		}

		fullPath := storageDir + logDir.Name()
		_, err = os.Stat(fullPath)
		if err == nil {
			err = os.Remove(fullPath)
			if err != nil {
				xtremelog.Error(err, false)
			}
		}

		fmt.Println(logDate.Format("2006-01-02 15:04:05"))
	}
}
