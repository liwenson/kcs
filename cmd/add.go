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
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加一个k8s 配置",
	Long:  `可以添加不同环境的配置，名称不能相同`,
	Run: func(cmd *cobra.Command, args []string) {

		// 参数校验

		if global.Name == "" {

			fmt.Println("必须指定名称")
			return
		}

		if global.Kubeconfig == "" {
			fmt.Println("必须指定配置文件")
			return
		}

		// 重名校验
		fileList, err := util.ListDir(global.KcsData)

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, name := range fileList {
			if name == global.Name {
				fmt.Println("改名称已经存在")
				return
			}
		}

		// 判断文件是否存在
		readFile, err := os.ReadFile(global.Kubeconfig)
		if err != nil {
			if os.IsNotExist(err) {
				// 文件不存在错误: 用于 只读打开 的时候
				fmt.Println("文件不存在:", err)
				return
			} else {
				fmt.Println(err)
				return
			}
		}

		err = os.WriteFile(fmt.Sprintf("%s/%s", global.KcsData, global.Name), readFile, 0666)

		if err != nil {
			fmt.Println(err)
			return
		}

		color.Green("%s,添加成功!", global.Name)

	},
}
