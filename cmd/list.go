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
				// 文件不存在错误: 用于只读打开 的时候
				fmt.Println("配置文件不存在:", err)
			} else {
				fmt.Println(err)
			}
			return
		}

		current := string(readFile)
		// 判断当前配置和标记的配置是否一致
		if verifyConfig(current) == false {
			color.Red("检验错误,请重新应用一个配置")
			current = ""

			// 不一致，则将config 清空，避免操作错误的集群
			// 创建一个空的字节切片
			emptyContent := []byte("")

			// 使用 ioutil.WriteFile 将空的字节切片写入文件
			err := os.WriteFile(fmt.Sprintf("%s/config", global.KubeHome), emptyContent, 0644)
			if err != nil {
				fmt.Println("Error clearing file:", err)
				return
			}
		}

		color.Green("当前: \t %s\n", current)
		for _, name := range fileList {
			color.Blue("\t %s", name)
		}
	},
}

func verifyConfig(current string) bool {

	currentFile, err := os.ReadFile(fmt.Sprintf("%s/%s", global.KcsData, current))
	if err != nil {
		if os.IsNotExist(err) {
			// 文件不存在错误: 用于 只读打开 的时候
			fmt.Println("文件不存在:", err)
		} else {
			fmt.Println(err)
		}
	}

	k8sFile, err := os.ReadFile(fmt.Sprintf("%s/config", global.KubeHome))
	if err != nil {
		if os.IsNotExist(err) {
			// 文件不存在错误: 用于 只读打开 的时候
			fmt.Println("文件不存在:", err)
		} else {
			fmt.Println(err)
		}
	}

	currentFileMd5 := util.MD5(currentFile)
	k8sFileMd5 := util.MD5(k8sFile)

	if currentFileMd5 == k8sFileMd5 {
		return true
	}

	return false

}
