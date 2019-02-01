package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yasukotelin/bcon/utf8"
)

var rootCmd = &cobra.Command{
	Use:   "bcon",
	Short: "count up byte number of string",
	Long:  "bcon is the tool to count up byte number of string",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			n := utf8.Count(args[0])
			fmt.Println(n)
		}
	},
}

// Execute はRootコマンドの実行関数
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
