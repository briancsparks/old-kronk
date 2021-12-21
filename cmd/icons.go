package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "io/ioutil"
  "path/filepath"

  "github.com/spf13/cobra"
)

// iconsCmd represents the icons command
var iconsCmd = &cobra.Command{
	Use:   "icons",
	Short: "Generate icons and graphics for projects.",
	Long: `Generate icons and graphics for projects.

Long`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("icons called")

    //file, err := ioutil.TempFile("", "prefix")
    //check(err)
    //defer os.Remove(file.Name())
    //defer file.Close()

    dir, err := ioutil.TempDir("", "kronk_icons")
    check(err)
    //defer os.RemoveAll(dir)

    fmt.Println(dir)

    convertArgs := []string{"-pointsize", "2400", "-fill", "royalblue", "-background", "none",
      "-flatten", "-font", "Courier-New", "label:*IO", "-trim", "+repage", filepath.Join(dir, "icon.png")}

    output, err := launch4Result("convert", convertArgs)
    check(err)

    <- output
    //fmt.Println(<- output)
	},
}

func init() {
	rootCmd.AddCommand(iconsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// iconsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// iconsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
