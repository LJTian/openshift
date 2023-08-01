package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "TCurl",
	Short: "TCurl is an http access command client",
	Long:  `TCurl is an http client mainly used for web service access and recording command line programs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("用法指导 TCurl uri http://127.0.0.1:8080/ping") // Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
