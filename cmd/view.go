package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"kcs/global"
	"os"
	"regexp"
)

// rootCmd represents the base command when called without any subcommands
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "查看配置",
	Long:  `查看kubernetes的config配置`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	// 使用自定义验证函数
	Args: func(cmd *cobra.Command, args []string) error {

		//if len(args) < 1 {
		//	return errors.New("必须指定一个配置")
		//}

		if len(args) > 1 {
			return errors.New("只需求指定一个配置")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		var cfgPath string

		if len(args) == 0 {
			cfgPath = fmt.Sprintf("%s/config", global.KubeHome)
		} else {
			cfgPath = fmt.Sprintf("%s/%s", global.KcsData, args[0])
		}

		configFile, err := os.ReadFile(cfgPath)
		if err != nil {
			if os.IsNotExist(err) {
				// 文件不存在错误: 用于 只读打开 的时候
				fmt.Println("文件不存在:", err)
			} else {
				fmt.Println(err)
			}
		}

		// 清屏
		//util.ClearScreen()

		if global.Server {
			server := getServer(string(configFile))
			if server != "" {
				fmt.Println(server)
			}
		} else {
			fmt.Println(string(configFile))
		}

	},
}

func getServer(content string) (server string) {
	pattern := `(server: https://.*)`

	// 编译正则表达式
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return ""
	}
	// 在文件内容中查找匹配项
	matches := regex.FindAllString(content, 1)
	// 打印匹配的字符串
	server = fmt.Sprintf("kubernetes %s", matches[0])
	return
}
