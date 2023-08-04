package cmd

import (
	"cobra-curl-cli/pkg/db"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all clientName",
	Long:  `list all clientName`,
	Run: func(cmd *cobra.Command, args []string) {
		if DbArgs.DbConnectUri == "" {
			fmt.Println("数据库相关内容为空")
			os.Exit(1)
		}
		if err := db.StartDB(DbArgs.DbConnectUri); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		names, err := db.ListClientName()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for x, name := range names {
			fmt.Printf("[%d]:[%s]\n", x, name)
		}
	},
}
