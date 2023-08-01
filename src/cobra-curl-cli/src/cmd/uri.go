package cmd

import (
	"cobra-curl-cli/pkg/envVar"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(uriCmd)
}

var uriCmd = &cobra.Command{
	Use:   "uri",
	Short: "uri service address",
	Long:  `uri service address`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		os.Setenv(envVar.Uri, args[0])
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		os.Getenv("")
	},
}
