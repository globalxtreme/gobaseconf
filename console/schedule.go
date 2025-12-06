package console

import (
	"github.com/globalxtreme/gobaseconf/console/command"
	"github.com/go-co-op/gocron"
	"time"
)

type callbackFunc func(*gocron.Scheduler)

func Schedules(callback callbackFunc) {
	sch := gocron.NewScheduler(time.UTC)

	// Schedules
	addSchedule(sch.Every(1).Day().At("00:01"), &command.DeleteLogFileCommand{FromSchedule: true})
	callback(sch)

	sch.StartBlocking()
}

func addSchedule(schedule *gocron.Scheduler, command BaseInterface) {
	schedule.Do(command.Handle)
}
