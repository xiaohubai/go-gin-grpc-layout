package scripts

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "scripts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root commands")
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&appConf, "app_conf", "", "app.yaml", "环境配置")

	cobra.CheckErr(rootCmd.Execute())
}

// 公用变量
var (
	appConf   string
	startDate string
)

func main() {
	Execute()
}
