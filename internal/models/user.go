package models

import "time"

type User struct {
	ID              uint
	Email           string
	Password        string
	EmailVerifiedAt *time.Time
}
