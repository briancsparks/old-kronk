package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "net"
  "os"
  "os/signal"
  "runtime"

  "github.com/spf13/cobra"
  "github.com/zserge/lorca"
)

// lorcaCmd represents the lorca command
var lorcaCmd = &cobra.Command{
	Use:   "lorca",
	Short: "An experiment using Lorca",
	Long: `An experiment using Lorca

Lorca, I say.`,

	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("lorca called")
    
    lorcaArgs := make([]string, 0)
    if runtime.GOOS == "linux" {
      lorcaArgs = append(lorcaArgs, "--class=Lorca")
    }

    ui, err := lorca.New("", "", 800, 600, lorcaArgs...)
    check(err)
    defer ui.Close()

    // connect to FS (fileServer pointing to folder www)
    listener, err := net.Listen("tcp", "127.0.0.1:0")
    check(err)
    defer listener.Close()

    // os signal handling
    sigc := make(chan os.Signal)
    signal.Notify(sigc, os.Interrupt)
    select {
    case <-sigc:
    case <-ui.Done():
    }
    // can exit now
    fmt.Println("Thanks for using the app!")
	},
}

func init() {
	experimentCmd.AddCommand(lorcaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lorcaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lorcaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
