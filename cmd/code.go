package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"

	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Do something with code / repos",
	Long: `Do something with code / repos

code is fun.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("code called")
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
