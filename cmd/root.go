package cmd

import (
	"fmt"
	"kcs/global"
	"kcs/util"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kvm",
	Short: "用来管理k8s配置",
	Long:  `通过命令的方式来管理多环境kubectl的配置的切换`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(useCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(viewCmd)

	addCmd.Flags().StringVarP(&global.Name, "name", "n", "", "config name")
	addCmd.Flags().StringVarP(&global.Kubeconfig, "kubeconfig", "c", "", "config path")
	viewCmd.Flags().BoolVarP(&global.Server, "service", "s", false, "查看k8s集群地址")

}

func initConfig() {
	// 初始话 kcs 目录
	_, err := util.PathExists(global.KcsHome)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = util.PathExists(global.KcsData)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = util.FileExists(global.KcsConfig, true)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = util.PathExists(fmt.Sprintf(global.KubeHome))
	if err != nil {
		fmt.Println(err)
		return
	}
}
