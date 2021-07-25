package models

import (
	"time"
)

type Alarm struct {
	Id        int64
	AlermTime *time.Time
	Content   string
	Status    int
}
