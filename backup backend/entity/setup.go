package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db
}
func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-64-3.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")
	}
	// Migrate the schema
	database.AutoMigrate(
		&Ambulance{},
		&User{},
		&CheckList{},
		&Car_path{},
		&Path_status{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&User{}).Create(&User{
		Name:     "Rattatammanoon",
		Email:    "rattatammanoontop@gmail.com",
		Password: string(password),
	})
	db.Model(&User{}).Create(&User{
		Name:     "Name",
		Email:    "name@example.com",
		Password: string(password),
	})
	var rattatammanoon User
	var name User
	db.Raw("SELECT * FROM users WHERE email = ?", "rattatammanoontop@gmail.com").Scan(&rattatammanoon)
	db.Raw("SELECT * FROM users WHERE email = ?", "name@example.com").Scan(&name)

	car1 := Ambulance{
		Brand:          "Nilsson Volvo SA1",
		Status:         "Ready to use",
		Ambulance_type: "AdvancedLifeSupport",
	}
	db.Model(&Ambulance{}).Create(&car1)

	car2 := Ambulance{
		Brand:          "Toyota Hilux Vigo SA2",
		Status:         "Ready to use",
		Ambulance_type: "AdvancedLifeSupport",
	}
	db.Model(&Ambulance{}).Create(&car2)

	car3 := Ambulance{
		Brand:          "Honda Stepwgn Spada SA3",
		Status:         "Ready to use",
		Ambulance_type: "AdvancedLifeSupport",
	}
	db.Model(&Ambulance{}).Create(&car3)

	status1 := Path_status{
		Status: "Normal",
	}
	db.Model(&Path_status{}).Create(&status1)

	status2 := Path_status{
		Status: "Defective/Fix",
	}
	db.Model(&Path_status{}).Create(&status2)

	status3 := Path_status{
		Status: "Order",
	}
	db.Model(&Path_status{}).Create(&status3)

	path1 := Car_path{
		Path_titel: "oil fuel",
	}
	db.Model(&Car_path{}).Create(&path1)
	path2 := Car_path{
		Path_titel: "oil fuel1",
	}
	db.Model(&Car_path{}).Create(&path2)
	path3 := Car_path{
		Path_titel: "oil fuel2",
	}
	db.Model(&Car_path{}).Create(&path3)

	db.Model(&CheckList{}).Create(&CheckList{
		Checked_time: time.Now(),
		Ambulance:    car1,
		Car_path:     path1,
		Path_status:  status1,
	})

}
