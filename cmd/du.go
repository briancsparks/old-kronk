
/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// duCmd represents the du command
var duCmd = &cobra.Command{
	Use:   "du",

	Short: "Compute disk usage.",
	Long: `Recursively visits all sub-directories, and calculates how much space any area is.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("du called")
	},
}

func init() {
	rootCmd.AddCommand(duCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// duCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// duCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
