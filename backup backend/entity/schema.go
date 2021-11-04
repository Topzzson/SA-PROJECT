package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CheckList []CheckList `gorm:"foreignKey:OwnerID"`
}

type Ambulance struct {
	gorm.Model
	Brand          string `gorm:"uniqueIndex"`
	Status         string
	Ambulance_type string
	CheckList      []CheckList `gorm:"foreignKey:AmbulanceID"`
}

type Car_path struct {
	gorm.Model

	Path_titel string `gorm:"uniqueIndex"`

	CheckList []CheckList `gorm:"foreignKey:Car_pathID"`
}
type Path_status struct {
	gorm.Model
	Status    string      `gorm:"uniqueIndex"`
	CheckList []CheckList `gorm:"foreignKey:Path_statusID"`
}

type CheckList struct {
	gorm.Model
	Checked_time time.Time
	OwnerID      *uint
	Owner        User

	Car_pathID *uint
	Car_path   Car_path

	Path_statusID *uint
	Path_status   Path_status

	AmbulanceID *uint
	Ambulance   Ambulance
}
