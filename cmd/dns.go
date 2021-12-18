package cmd

/* Copyright © 2021 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fqdn string
var domainName string
var ipAddress string
var recordType string

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Manage DNS",
	Long: `Manage DNS

at Route53.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dns called")
	},
}

func init() {
	rootCmd.AddCommand(dnsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dnsCmd.PersistentFlags().String("foo", "", "A help for foo")

  //dnsCmd.PersistentFlags().StringVarP(&fqdn, "fqdn", "f", "", "Fully qualified domain name")
  //dnsCmd.PersistentFlags().StringVarP(&domainName, "domain", "d", "", "Domain name")
  //dnsCmd.PersistentFlags().StringVarP(&ipAddress, "ip", "i", "", "IP address")
  //dnsCmd.PersistentFlags().StringVarP(&recordType, "type", "r", "A", "Record type")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
