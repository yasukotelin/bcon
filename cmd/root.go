package cmd

import (
	"fmt"
	"os"

	"github.com/yasukotelin/bcon/counter"

	"github.com/spf13/cobra"
)

const (
	v = "1.0.0"
)

var (
	cset string
)

var rootCmd = &cobra.Command{
	Use:     "bcon",
	Version: v,
	Short:   "count up byte number of string",
	Long:    "bcon is the tool to count up byte number of string",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			n, err := counter.Count(args[0], cset)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				fmt.Println(n)
			}
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&cset, "cset", "c", "utf8", "charset flag")
}

// Execute はRootコマンドの実行関数
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
