package goutils

import (
	"time"
)

type RetryConfig struct {
	// simple retry configs
	Attempts uint
	Interval time.Duration

	attempts uint
}

func Retry(fn func(), cfg *RetryConfig) {
	if cfg == nil {
		cfg = new(RetryConfig)
	}
	cfg.attempts++
	err := Try(fn)
	if err != nil {
		if cfg.attempts < cfg.Attempts {
			if cfg.Interval > 0 {
				time.Sleep(cfg.Interval)
			}
			Retry(fn, cfg)
		} else {
			panic(err)
		}
	}
}
