package model

import "gorm.io/gorm"

type session struct {
	*gorm.Model
	UserId string
}
