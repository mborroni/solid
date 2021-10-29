package dependency

import (
	"time"
)

const (
	interval = 5
)

type Cron struct {
	ticker   *time.Ticker
	interval time.Duration
	cleaner  *DBCleaner
}

func NewCron() *Cron {
	return &Cron{
		cleaner:  NewDBCleaner(),
		interval: interval * time.Minute,
	}
}

// Init executes an action every 5 minutes interval
func (c *Cron) Init() {
	c.ticker = time.NewTicker(c.interval)
	go func() {
		for {
			select {
			case <-c.ticker.C:
				if err := c.cleaner.Clean(); err != nil {
					c.ticker.Stop()
				}
			}
		}
	}()
}

// Stop stops the execution of the cron
func (c *Cron) Stop() {
	c.ticker.Stop()
}
