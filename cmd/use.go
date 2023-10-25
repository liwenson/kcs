package cmd

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"kcs/global"
	"kcs/util"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "切换配置",
	Long:  `切换到指定的配置`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	// 使用自定义验证函数
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

		config := fmt.Sprintf("%s/%s", global.KcsData, args[0])
		// 校验是否存在
		exists, err := util.FileExists(config, false)
		if err != nil {
			fmt.Println(err)
			return
		}

		if !exists {
			color.Red("%s,不存在!", args[0])
			return
		}

		// 判断文件是否存在
		readFile, err := os.ReadFile(config)
		if err != nil {
			if os.IsNotExist(err) {
				// 文件不存在错误: 用于 只读打开 的时候
				fmt.Println("文件不存在:", err)
			} else {
				fmt.Println(err)
			}
		}

		err = os.WriteFile(fmt.Sprintf("%s/config", global.KubeHome), readFile, 0666)
		if err != nil {
			fmt.Println(err)
		}

		// 保存当前的记录
		err = os.WriteFile(fmt.Sprintf("%s", global.KcsConfig), []byte(args[0]), 0666)
		if err != nil {
			fmt.Println(err)
		}

		color.Green("%s,切换成功!", args[0])
	},
}
