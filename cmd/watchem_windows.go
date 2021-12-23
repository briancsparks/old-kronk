package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "log"
)

// watchemCmd represents the watchem command
var watchemCmd = &cobra.Command{
	Use:   "watchem",

	Short: "Have Kronk watch the typical items.",
	Long: `Have Kronk watch the typical items.

For example, watch for when the proxy registry changes.`,

	Run: func(cmd *cobra.Command, args []string) {
    //fmt.Println("watchem called")

    isVpn, err := cmd.Flags().GetBool("vpn")
    if err != nil {
      log.Fatal(err)
    }

    isProxy, problem, err := checkProxy(isVpn)
    if err != nil {
      log.Fatal(err)
    }
    if len(problem) > 0 {
      fmt.Println(problem)
    }

    fmt.Printf("Proxy enabled: %v\n", isProxy)
  },

}

func init() {
	rootCmd.AddCommand(watchemCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
  // watchemCmd.PersistentFlags().String("foo", "", "A help for foo")
  watchemCmd.PersistentFlags().Bool("vpn", false, "Assume on VPN")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watchemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



/*
Copyright © 2021 Brian C Sparks <briancsparks@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
