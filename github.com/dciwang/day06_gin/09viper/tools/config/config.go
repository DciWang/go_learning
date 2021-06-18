package config

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	cfgDatabases *viper.Viper
)

//载入配置文件
func ConfigSetup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Printf(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	// 数据库初始化
	cfgDatabases = viper.Sub("settings.databases")
	if cfgDatabases != nil {
		panic("config not found settings.database")
	}

	databaseConfig = InintDatabase(cfgDatabases)

}
