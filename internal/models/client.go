package models

import "time"

type Client struct {
	ID       int
	PlayTime time.Duration
	Timeout  time.Duration
	Start    chan struct{}
}
