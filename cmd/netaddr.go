/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
)

// netaddrCmd represents the netaddr command
var netaddrCmd = &cobra.Command{
	Use:   "netaddr",
	Short: "returns the network address for a given CIDR address",
	Long: `This command returns the network address for a given CIDR address. e.g.:
  ipcalc netaddr 192.168.178.86/24`,
	Run: func(cmd *cobra.Command, args []string) {
		_, network, err := net.ParseCIDR(args[0])
		if err != nil {
			log.Panic(err)
			return
		}
		println(network.IP.String())
	},
}

func init() {
	rootCmd.AddCommand(netaddrCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netaddrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netaddrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
