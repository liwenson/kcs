package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"kcs/global"
	"kcs/util"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出当前可用的k8s配置",
	Long:  `列出当前可用的k8s配置`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

	Run: func(cmd *cobra.Command, args []string) {
		fileList, err := util.ListDir(global.KcsData)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 获取当前的配置

		readFile, err := os.ReadFile(fmt.Sprintf("%s", global.KcsConfig))
		if err != nil {
			if os.IsNotExist(err) {
				// 文件不存在错误: 用于 只读打开 的时候
				fmt.Println("文件不存在:", err)
			} else {
				fmt.Println(err)
			}
		}
		current := string(readFile)

		color.Green("当前: \t %s\n", current)
		for _, name := range fileList {
			color.Blue("\t %s", name)

		}
	},
}
