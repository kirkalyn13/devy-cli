package utils

import (
	"log"
	"time"

	cron "github.com/robfig/cron/v3"
)

const RUNTIME = 8 //8 hrs job

// ========== Run Cron Job ==========
func RunJob(task func(), duration string) {
	cron := cron.New()
	log.Print("Created Cron Job.")

	cron.AddFunc(duration, func() {
		task()
	})

	log.Print("Starting Cron Job.")
	cron.Start()
	time.Sleep(time.Hour * RUNTIME)
}
