package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var SynthCmd = cobra.Command{
	Use:     "synth <IPv6 prefix subnet> <IPv4 address>",
	Short:   "Synthesis of IPv6 address from IPv4 according to RFC6147",
	Example: "synth6 synth 64:ff9b::/96 192.0.0.171",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		addr := net.ParseIP(args[1])
		if addr == nil {
			fmt.Printf("provided address %s is not IP\n", args[1])
			os.Exit(1)
		}
		addr = addr.To4()
		if addr == nil {
			fmt.Printf("provided address %s is not IP\n", args[1])
			os.Exit(1)
		}

		_, prefix, err := net.ParseCIDR(args[0])
		if err != nil {
			fmt.Printf("failed to parse net %s\n", args[0])
			os.Exit(1)
		}

		n, _ := prefix.Mask.Size()
		// Assumes prefix has been validated during setup
		v6 := make([]byte, 16)
		i, j := 0, 0

		for ; i < n/8; i++ {
			v6[i] = prefix.IP[i]
		}
		for ; i < 8; i, j = i+1, j+1 {
			v6[i] = addr[j]
		}
		if i == 8 {
			i++
		}
		for ; j < 4; i, j = i+1, j+1 {
			v6[i] = addr[j]
		}

		fmt.Println(net.IPNet{IP: v6}.IP)

	},
}

func init() {
	RootCmd.AddCommand(&SynthCmd)
}
