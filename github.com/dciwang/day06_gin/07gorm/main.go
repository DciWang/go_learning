package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type PmsBrand struct {
	BrandId     int    `json:"brand_id" grom:"brand_id"`
	name        string `json:"name" grom:"name"`
	Logo        string `json:"logo" grom:"logo"`
	ShowStatus  int    `json:"show_status" gorm:"show_status"`
	FirstLetter string `json:"first_letter" grom:"first_letter"`
	Sort        int    `json:"sort" grom:"sort"`
}

func main() {
	r := gin.Default()

	db, err := gorm.Open("mysql", "root:dev@(60.205.194.87:3306)/gulimall_pms?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("failed")
		panic(err)
	}
	fmt.Println("success")
	fmt.Printf("%#v\n", db)
	defer db.Close()
	var brand PmsBrand
	db.AutoMigrate(brand)
	var brands []PmsBrand
	brand.name ="华为"
	if err := db.Find(&brands).Error; err != nil {
		panic(err)
	}

	for i := range brands {
		fmt.Println("i---->",i)
	}
	fmt.Println("brands --->", brand)
	_ = r.Run(":9090")
}
