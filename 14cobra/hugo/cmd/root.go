package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// cobra.Command代表一个命令
// Use 命令的名称
// Short 命令的简单描述
// Long 命令的完整描述
// Run 执行命令的时候调用的函数

// rootCmd.Execute()是命令的执行入口
// 内部会解析os.Args[1:]参数列表(默认情况下是这样，也可以通过Command.SetArgs方法设置参数)
// 然后遍历命令树，为命令找到合适的匹配项对应的标志

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "hugo is a very fast static site generator",
	Long: `A Fast and Fleible Static Site Generator built wirh love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run hugo")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
