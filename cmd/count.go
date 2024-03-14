package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"

	"github.com/jmbit/ipcalc/maths"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "counts valid host IPs in a subnet (CIDR notation)",
	Long: `This command counts the number of valid host IPs in a subnet. e.g.:
  ipcalc count 192.168.178.0/24.
  This does not substract a possible gateway address`,
	Run: func(cmd *cobra.Command, args []string) {
		_, network, err := net.ParseCIDR(args[0])
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Printf("%d\n", maths.AddressCount(network))

	},
}

func init() {
	rootCmd.AddCommand(countCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
