package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"kvm/global"
	"kvm/util"
)


// rootCmd represents the base command when called without any subcommands
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出当前可用的k8s配置",
	Long: `列出当前可用的k8s配置`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(global.KvmHome)

		fileList, err := util.ListDir(global.KvmHome)
		if err != nil {
			fmt.Println(err)
			return
		}


		for _, name := range fileList {
			//fmt.Println(info.Name())
			color.Blue(name)
		}
	},
}
