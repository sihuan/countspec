package mysql

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

// func init() {
// 	dsn := "mysql:@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	DB = db
// 	DB.AutoMigrate(&models.Tarball{})

// }

func CheckDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

}
