package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "github.com/spf13/cobra"
  "path/filepath"
)

var zero bool

// generateLorcaCmd represents the generate command
var generateLorcaCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Lorca assets",
	Long: `Generate Lorca assets.

Lorca assets`,

	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("generate called")

    myDirname := MyDirName()
    outFilename := filepath.Join(myDirname, "lorcawebassets.go")

    if ! zero {
      assetSrc := filepath.Join("D:", "data", "projects", "WebStormProjects", "reactjs", "two", "for-lorca-redux", "build")
      fmt.Printf("Generating lorca assets from %s\n", assetSrc)
      err := KronkEmbed("cmd", "FS", outFilename, assetSrc)
      Check(err)

    } else {
      err := KronkEmbed0("cmd", "FS", outFilename)
      Check(err)
    }

	},
}

func init() {
	lorcaCmd.AddCommand(generateLorcaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateLorcaCmd.PersistentFlags().String("foo", "", "A help for foo")
  generateLorcaCmd.PersistentFlags().BoolVarP(&zero, "zero", "z", false, "Use a zero-value")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateLorcaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
