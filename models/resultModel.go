package models

import (
    "time"
    "gorm.io/gorm"
)

type Result struct {
    gorm.Model
    MonitorID      uint       `gorm:"index;not null" json:"monitor_id"`
    Status         string     `gorm:"size:50" json:"status_code"`
    CheckedAt      time.Time  `json:"checked_at"`                
    ResponseMs     int64      `json:"response_ms"`               
    Details        string     `gorm:"size:500" json:"details"`
}
