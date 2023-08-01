package cmd

import (
	"cobra-curl-cli/pkg/curl"
	"cobra-curl-cli/pkg/db"
	"cobra-curl-cli/pkg/define"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var AllArgs define.TCurl
var DbArgs define.DBInfo

func init() {
	rootCmd.PersistentFlags().StringVarP(&AllArgs.Uri, "uri", "U", "http://127.0.0.1:8080/ping", "web service addr")
	rootCmd.PersistentFlags().IntVarP(&AllArgs.Times, "times", "T", 100, "Number of cycles")
	rootCmd.PersistentFlags().IntVarP(&AllArgs.Intervals, "intervals", "I", 5, "Intervals")
	rootCmd.PersistentFlags().IntVarP(&AllArgs.TimeOut, "timeout", "t", 5, "timeout period")
	rootCmd.PersistentFlags().BoolVarP(&AllArgs.SaveDB, "saveDB", "S", false, "Whether the data is saved in the database")

	rootCmd.PersistentFlags().StringVarP(&DbArgs.DbConnectUri, "dbUri", "D", "", "Database connect address")
}

var rootCmd = &cobra.Command{
	Use:   "TCurl",
	Short: "TCurl is an http access command client",
	Long:  `TCurl is an http client mainly used for web service access and recording command line programs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(AllArgs)
		if AllArgs.SaveDB {
			if DbArgs.DbConnectUri == "" {
				fmt.Println("数据库相关内容为空")
				os.Exit(1)
			}
			if err := db.StartDB(DbArgs.DbConnectUri); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		if err := curl.Run(AllArgs, DbArgs); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
