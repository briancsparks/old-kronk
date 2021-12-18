package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
  "github.com/spf13/viper"

  "github.com/spf13/cobra"
)

// setrecordCmd represents the setrecord command
var setrecordCmd = &cobra.Command{
	Use:   "setrecord",
	Short: "Set a DNS record.",
	Long: `Set a DNS record.

Sets a DNS record at Route53`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setrecord called")
    fmt.Printf("fqdn: %s\n", fqdn)
    fmt.Printf("dn: %s\n", domainName)
    fmt.Printf("IP: %s\n", ipAddress)
    fmt.Printf("Record Type: %s\n", recordType)
	},
}

func init() {
	dnsCmd.AddCommand(setrecordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setrecordCmd.PersistentFlags().String("foo", "", "A help for foo")

  setrecordCmd.PersistentFlags().StringVarP(&fqdn, "fqdn", "f", "", "Fully qualified domain name")
  setrecordCmd.PersistentFlags().StringVarP(&ipAddress, "ip", "i", "", "IP address (required)")
  setrecordCmd.PersistentFlags().StringVarP(&recordType, "type", "r", "A", "Record type")

  setrecordCmd.PersistentFlags().StringVarP(&domainName, "domain", "d", "", "Domain name")
  viper.BindPFlag("domain", rootCmd.PersistentFlags().Lookup("domain"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setrecordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
