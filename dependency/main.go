package main

import (
	"time"
)

func main() {
	interval := 5 * time.Second
	cleaner := NewDBCleaner()

	// libertad para mockear
	cron := NewCron(interval, cleaner)
	cron.Init()
}