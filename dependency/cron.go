package main

import (
	"time"
)

type Actionable interface {
	Do() error
}

type Cron struct {
	ticker   *time.Ticker
	interval time.Duration
	action   Actionable
}

func NewCron(interval time.Duration, actionable Actionable) *Cron {
	return &Cron{
		action:   actionable,
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
				if err := c.action.Do(); err != nil {
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
