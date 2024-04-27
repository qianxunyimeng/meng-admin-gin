// @Author [shiqingliang](https://github.com/qianxunyimeng)
// @Date 2024/4/27 21:42:00
// @Desc

package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"meng-admin-gin/cmd/api"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "meng-admin",       // 命令名称
	Short:             "meng-admin 命令行工具", // 命令描述
	SilenceUsage:      true,               // 如果命令执行出错，它会阻止打印用法信息
	DisableAutoGenTag: true,               // 它会禁止在生成的帮助信息中添加日期标签
	//TraverseChildren:  false,              //
	Long: `meng-admin 命令行工具`, // 详细描述
	Args: func(cmd *cobra.Command, args []string) error { // 用于验证传递给命令的参数
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil }, // PersistentPreRunE 字段定义了一个函数，它会在命令的任何子命令执行之前运行
	Run: func(cmd *cobra.Command, args []string) { // 这个函数是命令的实际执行体
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 meng-admin 0.0.1 可以使用 -h 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
