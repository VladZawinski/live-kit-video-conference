package model

import "time"

type Room struct {
	ID          int
	Name        string
	Description string
	OwnerID     int
	CreatedAt   time.Time
}
