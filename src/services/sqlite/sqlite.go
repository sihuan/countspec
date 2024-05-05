package sqlite

import (
	"mygo/core/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MyGODB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("data/mygo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Tarball{}, &models.Task{}, &models.QemuTask{})
	MyGODB = db

}

func CheckDB() {
	sqlDB, err := MyGODB.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}
}
