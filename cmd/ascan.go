package cmd

import (

	"github.com/spf13/cobra"
	zap "github.com/hahwul/mzap/pkg/zap"
)

// ascanCmd represents the ascan command
var ascanCmd = &cobra.Command{
	Use:   "ascan",
	Short: "Add ActiveScan ZAP",
	Run: func(cmd *cobra.Command, args []string) {
		zap.ActiveScan(URLs,apiHosts)
	},
}

func init() {
	rootCmd.AddCommand(ascanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ascanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ascanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
