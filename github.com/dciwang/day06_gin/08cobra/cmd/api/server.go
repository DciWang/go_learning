package api

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	config   string
	port     string
	mode     string
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start API server",
		Example: "dc server config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			hello()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func hello() {
	fmt.Println("hello dc")
}
func run() error {
	fmt.Printf("prot is %s\n", port)
	return nil
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8002", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}
