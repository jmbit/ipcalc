package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"

	"github.com/jmbit/ipcalc/maths"
)

// firstCmd represents the first command
var firstCmd = &cobra.Command{
	Use:   "first",
	Short: "Print the first IP address in a subnet (CIDR notation)",
	Long: `Prints the first IP address in a subnet. e.g.:
  ipcalc first 192.168.170.0/24
  `,

	Run: func(cmd *cobra.Command, args []string) {
		_, network, err := net.ParseCIDR(args[0])
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Print(maths.GetFirstIPAddress(network).String() + "\n")

	},
}

func init() {
	rootCmd.AddCommand(firstCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// firstCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// firstCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
