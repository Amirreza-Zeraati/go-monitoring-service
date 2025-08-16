package services

import (
	"go-monitoring-service/initializers"
	"go-monitoring-service/models"
    "net/http"
	"time"
	"fmt"
)

func MonitorScheduler() {
    go func() {
        for {
            var monitors []models.Monitor
            initializers.DB.Where("active = ?", true).Find(&monitors)
            now := time.Now()
            for _, monitor := range monitors {
                if monitor.LastCheckedAt.IsZero() || now.Sub(monitor.LastCheckedAt) >= monitor.Interval {
                    go MonitorCheck(monitor)
                }
            }
            time.Sleep(1 * time.Minute)
        }
    }()
}

func MonitorCheck(monitor models.Monitor) {
    fmt.Printf("Checking monitor: %s at %v\n", monitor.Name, time.Now())
    
    var status string
	var details string
	var latency int64

	start := time.Now()

	switch monitor.Type {
	case "http":
		resp, err := http.Get(monitor.Target)
		if err != nil {
			status = "DOWN"
			details = err.Error()
		} else {
			defer resp.Body.Close()
			latency = time.Since(start).Milliseconds()

			if resp.StatusCode == monitor.ExpectedStatus {
				status = "UP"
			} else {
				status = "DOWN"
				details = fmt.Sprintf("Expected %d but got %d", monitor.ExpectedStatus, resp.StatusCode)
			}
		}

	case "dns/ssl":
		// TODO
		status = "NOT_IMPLEMENTED"
		details = "DNS/SSL check not implemented yet"

	case "db":
		// TODO
		status = "NOT_IMPLEMENTED"
		details = "DB check not implemented yet"

	case "deadman":
		// TODO
		status = "NOT_IMPLEMENTED"
		details = "Dead Man's Switch check not implemented yet"

	default:
		status = "UNKNOWN"
		details = "Unsupported monitor type"
	}

	result := models.Result{
		MonitorID:    monitor.ID,
		Status:       status,
		ResponseMs:   latency,
		CheckedAt:    time.Now(),
		Details:      details,
	}

	if err := initializers.DB.Create(&result).Error; err != nil {
		fmt.Println("Error saving result:", err)
	}

    monitor.LastCheckedAt = time.Now()
    monitor.LastStatus = status
    monitor.LastResponseMs = latency
    initializers.DB.Save(&monitor)
}
