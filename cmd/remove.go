package cmd

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"kcs/global"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "删除一个k8s 配置",
	Long:  `删除不同环境的配置`,
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return errors.New("必须指定一个配置")
		}

		if len(args) != 1 {
			return errors.New("只需求指定一个配置")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		// 判断文件是否存在

		err := os.Remove(fmt.Sprintf("%s/%s", global.KcsData, args[0]))
		if err != nil {
			if os.IsNotExist(err) {
				// 文件不存在错误: 用于 只读打开 的时候
				fmt.Println("文件不存在:", err)
			} else {
				fmt.Println(err)
			}
		}

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
		if args[0] == current {
			// 保存当前的记录
			err = os.WriteFile(fmt.Sprintf("%s", global.KcsConfig), []byte(""), 0666)
			if err != nil {
				fmt.Println(err)
			}
		}

		color.Green("删除 %s 成功!", args[0])
	},
}
