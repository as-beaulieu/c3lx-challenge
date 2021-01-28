package service

import (
	"c3lx-challenge/src/models"
	"time"
)

type HealthReporter interface {
	Heartbeat() (*models.HeartbeatResponse, error)
}

func (s service) Heartbeat() (*models.HeartbeatResponse, error) {
	return &models.HeartbeatResponse{
		Message: "API Running",
		Time:    time.Now(),
	}, nil
}
