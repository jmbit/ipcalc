package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"

	"github.com/jmbit/ipcalc/maths"
)

// lastCmd represents the last command
var lastCmd = &cobra.Command{
	Use:   "last",
	Short: "Print the last IP address in a subnet (CIDR notation)",
	Long: `Print the last IP address in a subnet. e.g.:
  ipcalc last 192.168.178.0/24
  `,
	Run: func(cmd *cobra.Command, args []string) {
		_, network, err := net.ParseCIDR(args[0])
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Print(maths.GetLastIPAddress(network).String() + "\n")

	},
}

func init() {
	rootCmd.AddCommand(lastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
