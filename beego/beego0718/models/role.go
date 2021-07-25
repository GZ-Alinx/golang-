package models

import (
	"time"
)

type Role struct {
	Id        int64
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeleteAt  *time.Time
}
