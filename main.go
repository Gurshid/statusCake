package main

import (
	"./Config"
	"./Models"
	"./Routers"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var err error

func initialMigration() {

	Config.DB, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("status: ", err)
	}
	//defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.UrlMaster{})
}

func main() {

	initialMigration()

	Routers.SetupRouter().Run()

}

