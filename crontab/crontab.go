package crontab

import (
	"github.com/robfig/cron"
)

func Go() {
	crontab := cron.New()

	crontab.Start()
}
