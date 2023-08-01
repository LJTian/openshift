package cmd

import (
	"cobra-curl-cli/pkg/envVar"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	uriCmd.AddCommand(timesCmd)
}

var timesCmd = &cobra.Command{
	Use:   "times",
	Short: "times is number of visits",
	Long:  `times is number of visits`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Setenv(envVar.Times, args[0])
	},
}
