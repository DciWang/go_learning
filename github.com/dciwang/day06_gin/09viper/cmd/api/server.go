package api

import (
	"fmt"
	"github.com/spf13/cobra"
	config2 "go_learning/github.com/dciwang/day06_gin/09viper/tools/config"
)

var (
	config   string
	port     string
	mode     string
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start Api server",
		Example: "viper server -c config/settings.dev.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("初始化")
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("初始化1")
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8002", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}

func run() error {
	fmt.Println("初始化  run")
	return nil
}

func setup() {
	fmt.Println("初始化系统配置")
	config2.ConfigSetup(config)
	fmt.Println("初始化数据库配置")
	config2.DBInit()
}
