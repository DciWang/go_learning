package cmd

import (
	"github.com/spf13/cobra"
	"go_learning/github.com/dciwang/day06_gin/09viper/cmd/api"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "viper", //mast   command
	Short:             "-v",    //short description
	SilenceUsage:      true,    //  be silence  when have error
	DisableAutoGenTag: true,    // 将通过为此命令生成文档来打印。
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `欢迎使用 viper，可以使用 -h 查看命令`
		log.Printf("%s\n", usageStr)
	},
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}

}
