package cmd

import (
	"fmt"
	"kvm/global"
	"kvm/util"
	"os"

	"github.com/spf13/cobra"
)




// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kvm",
	Short: "用来管理k8s配置",
	Long: `通过命令的方式来管理多环境kubectl的配置的切换`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.proj.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	cobra.OnInitialize(initConfig)
	
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(useCmd)

	addCmd.Flags().StringVarP(&global.Name, "name", "n", "", "config name")
	addCmd.Flags().StringVarP(&global.Kubeconfig, "kubeconfig", "c", "", "config path")

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func  initConfig()  {

	_, err := util.PathExists(fmt.Sprintf("%s/.kvm",global.Home))
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