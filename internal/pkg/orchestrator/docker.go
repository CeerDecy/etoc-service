package orchestrator

import (
	"gorm.io/gorm"

	"etoc-service/internal/app/http/db"
)

const (
	DockerAccessModeLocalEnv   = "LOCAL"
	DockerAccessModeTCP        = "TCP"
	DockerAccessModeTCPWithTLS = "TCP/TLS"
	DockerAccessModeAgent      = "AGENT"
)

type DockerInfo struct {
	Name       string `json:"name"`
	IP         string `json:"IP"`
	Port       string `json:"port"`
	AccessMode string `json:"accessMode"`
}

type Docker struct {
	db *db.DockerDB
}

func NewDocker(d *gorm.DB) *Docker {
	return &Docker{db: db.NewDockerDB(d)}
}
