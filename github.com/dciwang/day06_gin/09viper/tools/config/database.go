package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
)

type Database struct {
	Dbtype   string
	Host     string
	Port     int
	Dbname   string
	Username string
	Password string
}

var databaseConfig = &Database{}

var (
	err error
	DB  *gorm.DB
)

func InintDatabase(cfg *viper.Viper) *Database {
	return &Database{
		Port:     cfg.GetInt("port"),
		Dbtype:   cfg.GetString("dbType"),
		Host:     cfg.GetString("host"),
		Dbname:   cfg.GetString("Dbname"),
		Username: cfg.GetString("username"),
		Password: cfg.GetString("password"),
	}
}
func DBInit() {
	fmt.Println("数据库 init")
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Dbname,
		true,
		"Local")
	DB, err = gorm.Open(databaseConfig.Dbtype, config)
	if err != nil {
		panic(err)
	}
	DB.SingularTable(true)
	fmt.Println("数据库 init 结束...")
}
