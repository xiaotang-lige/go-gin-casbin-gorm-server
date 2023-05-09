package model

import "time"

type Authentication struct {
	UserId     string
	Token      string
	CreateTime time.Time
	Chanl      int
}
