package models

import (
	"gorm.io/gorm"
	"time"
)

type Monitor struct {
	gorm.Model
	UserID         uint          `gorm:"index;not null" json:"user_id"`
	Name           string        `gorm:"size:100;not null" json:"name"`
	Type           string        `gorm:"size:50;not null" json:"type"`
	Target         string        `gorm:"size:500" json:"target"`
	Method         string        `gorm:"size:10" json:"method"`
	ExpectedStatus int           `json:"expected_status"`
	Keyword        string        `gorm:"size:255" json:"keyword"`
	Interval       time.Duration `json:"interval"`
	Retries        int           `json:"retries"`
	Config         string        `gorm:"type:jsonb" json:"config"`
	Active         bool          `gorm:"default:true" json:"active"`
}
