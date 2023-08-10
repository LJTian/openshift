package cmd

import (
	"cobra-curl-cli/pkg/curl"
	"cobra-curl-cli/pkg/db"
	"cobra-curl-cli/pkg/define"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var AllArgs define.TCurl
var DbArgs define.DBInfo

func init() {
	rootCmd.PersistentFlags().StringVarP(&AllArgs.Uri, "uri", "U", "", "web service addr")
	rootCmd.PersistentFlags().StringVarP(&AllArgs.ClientName, "clientName", "N", "", "client name")
	rootCmd.PersistentFlags().IntVarP(&AllArgs.Times, "times", "T", 100, "Number of cycles")
	rootCmd.PersistentFlags().IntVarP(&AllArgs.Intervals, "intervals", "I", 5, "Intervals")
	rootCmd.PersistentFlags().IntVarP(&AllArgs.TimeOut, "timeout", "t", 5, "timeout period")
	rootCmd.PersistentFlags().IntVarP(&AllArgs.CoroutineNum, "coroutineNum", "c", 5, "coroutine number")
	rootCmd.PersistentFlags().BoolVarP(&AllArgs.SaveDB, "saveDB", "S", false, "Whether the data is saved in the database")

	rootCmd.PersistentFlags().StringVarP(&DbArgs.DbConnectUri, "dbUri", "D", "", "Database connect address")
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 生成随机字符串
func generateRandomString(length int) string {

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

var rootCmd = &cobra.Command{
	Use:   "TCurl",
	Short: "TCurl is an http access command client",
	Long:  `TCurl is an http client mainly used for web service access and recording command line programs`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(AllArgs)
		if AllArgs.Uri == "" {
			fmt.Println("URI 为空， 请使用 -h 查看使用说明")
			os.Exit(1)
		}
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

		fmt.Println("开始执行")
		if AllArgs.ClientName == "" {
			AllArgs.ClientName = fmt.Sprintf("clientName_%s", generateRandomString(5))
			fmt.Printf(" ClientName is [%s]\n", AllArgs.ClientName)
		}

		ch1 := make(chan int, AllArgs.CoroutineNum)
		for i := 0; i < AllArgs.CoroutineNum; i++ {
			go func(i int, ch chan int) {
				if err := curl.Run(AllArgs, i); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				ch1 <- 1
			}(i, ch1)
		}
		for i := 0; i < AllArgs.CoroutineNum; i++ {
			<-ch1
		}
		fmt.Println("循环结束")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
