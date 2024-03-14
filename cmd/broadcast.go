/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"

	"github.com/jmbit/ipcalc/maths"
)

// broadcastCmd represents the broadcast command
var broadcastCmd = &cobra.Command{
	Use:   "broadcast",
	Short: "Get the Broadcast address for a CIDR address",
	Long: `this command returns the broadcast address for a given CIDR address. e.g.:
  ipcalc broadcast 192.68.178.34/24
  `,
	Run: func(cmd *cobra.Command, args []string) {
		_, network, err := net.ParseCIDR(args[0])
		if err != nil {
			log.Panic(err)
			return
		}
		if ipaddr, err := maths.GetBroadcastAddress(network); err != nil {
			fmt.Print(err, "\n")
		} else {
			fmt.Print(ipaddr.String(), "\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(broadcastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// broadcastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// broadcastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
