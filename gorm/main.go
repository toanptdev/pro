package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Hooman struct {
	ID        int `json:"id"`
	DeletedAt gorm.DeletedAt
	Name      string `json:"name"`
	Age       int    `json:"age"`
}

type HoomanUpdate struct {
	Name *string `json:"name"`
	Age  int     `json:"age"`
}

func (Hooman) TableName() string {
	return "hooman"
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	// hooman := Hooman{
	// 	Name: "toan",
	// 	Age:  26,
	// }
	//
	// db.Create(&hooman)

	// var hooman Hooman
	// if err := db.Where("id = ?", 7).First(&hooman).Error; err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(hooman)
	//
	// var hoomans []Hooman
	// if err := db.Find(&hoomans).Error; err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println(hoomans)
	newName := ""

	if err := db.Table(Hooman{}.TableName()).Where("id = ?", 7).Updates(&HoomanUpdate{Name: &newName}).Error; err != nil {
		log.Fatal(err)
	}
}
