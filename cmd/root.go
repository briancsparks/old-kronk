package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var Verbose bool
var VVerbose bool
var VVVerbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kronk",

	Short: "The path that rocks",
	Long: `Hey, did you see that sky today -- Talk about blue!`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kronk.yaml)")

  rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
  viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

  rootCmd.PersistentFlags().BoolVarP(&VVerbose, "vverbose", "", false, "vverbose output")
  viper.BindPFlag("vverbose", rootCmd.PersistentFlags().Lookup("vverbose"))

  rootCmd.PersistentFlags().BoolVarP(&VVVerbose, "vvverbose", "", false, "vvverbose output")
  viper.BindPFlag("vvverbose", rootCmd.PersistentFlags().Lookup("vvverbose"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kronk" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kronk")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func rootInitForSub() {
  if VVVerbose {
    VVerbose = true
    Verbose  = true
  }

  if VVerbose {
    Verbose  = true
  }
}
