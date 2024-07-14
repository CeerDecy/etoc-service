package db

import (
	"gorm.io/gorm"

	"etoc-service/internal/app/http/db/model"
)

type DockerDB struct {
	db *gorm.DB
}

func NewDockerDB(DB *gorm.DB) *DockerDB {
	DB = DB.Model(model.Docker{})
	return &DockerDB{db: DB}
}

func (d *DockerDB) CreateDocker(docker model.Docker) error {
	return d.db.Create(docker).Error
}
