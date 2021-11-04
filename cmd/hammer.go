package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hammerCmd represents the hammer command
var hammerCmd = &cobra.Command{
	Use:   "hammer",

	Short: "If a hammer doesn't work, get a bigger hammer.",
	Long: `Ah, how shall I do it?
Oh, I know.
I'll turn him into a flea, a harmless, little flea,
and then I'll put that flea in a box,
and then I'll put that box inside of another box,
and then I'll mail that box to myself, and when it arrives...

...I'll smash it with a hammer!

It's brilliant, brilliant, brilliant, I tell you!
Genius, I say!
--------------------------------------------------------------------`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("If a hammer doesn't work, get a bigger hammer.")
	},
}

func init() {
	rootCmd.AddCommand(hammerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hammerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hammerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
