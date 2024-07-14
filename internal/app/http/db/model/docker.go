package model

import "gorm.io/gorm"

type Docker struct {
	gorm.Model
	Name       string `gorm:"column:name;type:varchar(255);not null"`
	IP         string `gorm:"column:ip;type:varchar(255)"`
	Port       string `gorm:"column:port;type:varchar(255)"`
	AccessMode string `gorm:"column:access_mode;type:varchar(255)"`
}

func (d *Docker) TableName() string {
	return "docker"
}
