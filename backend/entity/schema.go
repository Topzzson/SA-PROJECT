package entity

import (
	"time"

	"gorm.io/gorm"
)

type Informer struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Tel      string
	Password string

	Cases []Case `gorm:"foreignKey:InformerID"`
}

type Patient struct {
	gorm.Model
	Name   string
	Age    string
	Gender string

	Cases []Case `gorm:"foreignKey:PatientID"`
}

type Level struct {
	gorm.Model
	Rating string

	Cases []Case `gorm:"foreignKey:LevelID"`
}

type Characteristic struct {
	gorm.Model
	Category string

	Cases []Case `gorm:"foreignKey:CharacteristicID"`
}

type Case struct {
	gorm.Model
	CaseTime time.Time
	Address  string

	CharacteristicID *uint
	Characteristic   Characteristic

	LevelID *uint
	Level   Level

	InformerID *uint
	Informer   Informer

	PatientID *uint
	Patient   Patient

	Register []Register `gorm:"foreignKey:CaseID"`

	AssessmentSheet []AssessmentSheet `gorm:"foreignKey:CaseID"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Tel      string
	// 1 user ใช้ได้หลาย ambulance
	Ambulances []Ambulance `gorm:"foreignKey:RecorderID"`
	// 1 user ใช้ได้หลาย AssessmentSheets
	AssessmentSheets []AssessmentSheet `gorm:"foreignKey:RecorderID"`
}

type Operator struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Tel      string
	// 1 Operator ใช้ได้หลาย CheckList
	CheckLists []CheckList `gorm:"foreignKey:OperatorID"`
}

//-- Officer
type Officer struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Tel      string
	// 1 officer ใช้ได้หลาย Register
	Registers []Register `gorm:"foreignKey:OfficerID"`
}

type Register struct {
	gorm.Model
	RegisterTime time.Time

	// AssessmentSheet ทำหน้าที่เป็น FK
	AssessmentSheetID *uint
	AssessmentSheet   AssessmentSheet

	// Case ทำหน้าที่เป็น FK
	CaseID *uint
	Case   Case

	// Ambulance ทำหน้าที่เป็น FK
	AmbulanceID *uint
	Ambulance   Ambulance

	OfficerID *uint
	Officer   Officer
}

type Symptom struct {
	gorm.Model
	SymptomData string

	AssessmentSheet []AssessmentSheet `gorm:"foreignKey:SymptomID"`
}

type State struct {
	gorm.Model
	StateData string

	AssessmentSheet []AssessmentSheet `gorm:"foreignKey:StateID"`
}

type Assess struct {
	gorm.Model
	AssessData      string
	AssessmentSheet []AssessmentSheet `gorm:"foreignKey:AssessID"`
}

type AssessmentSheet struct {
	gorm.Model
	AssessTime time.Time

	CaseID *uint
	Case   Case

	SymptomID *uint
	Symptom   Symptom

	StateID *uint
	State   State

	AssessID *uint
	Assess   Assess

	RecorderID *uint
	Recorder   User

	Register []Register `gorm:"foreignKey:AssessmentSheetID"`
}

type Car_path struct {
	gorm.Model
	Path_titel string `gorm:"uniqueIndex"`

	CheckList []CheckList `gorm:"foreignKey:Car_pathID"`
}
type Path_status struct {
	gorm.Model
	Status string `gorm:"uniqueIndex"`

	CheckList []CheckList `gorm:"foreignKey:Path_statusID"`
}

type CheckList struct {
	gorm.Model
	Checked_time time.Time

	Car_pathID *uint
	Car_path   Car_path

	Path_statusID *uint
	Path_status   Path_status

	AmbulanceID *uint
	Ambulance   Ambulance

	OperatorID *uint
	Operator   Operator
}

type Ambulance struct {
	gorm.Model
	Registration  string
	RecordingTime time.Time

	AmbulanceTypeID *uint
	AmbulanceType   AmbulanceType

	StatusID *uint
	Status   Status

	BrandID *uint
	Brand   Brand

	RecorderID *uint
	Recorder   User

	Register  []Register  `gorm:"foreignKey:AmbulanceID"`
	CheckList []CheckList `gorm:"foreignKey:AmbulanceID"`
}
type AmbulanceType struct {
	gorm.Model
	TypeName   string
	Ambulances []Ambulance `gorm:"foreignKey:AmbulanceTypeID"`
}
type Status struct {
	gorm.Model
	StatusName string
	Ambulances []Ambulance `gorm:"foreignKey:StatusID"`
}
type Brand struct {
	gorm.Model
	BrandName  string
	Ambulances []Ambulance `gorm:"foreignKey:BrandID"`
}
