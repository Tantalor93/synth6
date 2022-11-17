package cmd

import (
	"context"
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var GetPrefix = cobra.Command{
	Use:     "prefix [IPv6 address from response of ipv4only.arpa]",
	Short:   "IPv6 prefix discovery for IPv6 address synthesis either using RFC7050 or based on already provided IPv6 address from ipv4only.arpa response",
	Example: "synth6 prefix 64:ff9b::c000:ab",
	Run: func(cmd *cobra.Command, args []string) {
		var ip net.IP
		if len(args) > 0 {
			ip = net.ParseIP(args[0])
			if ip == nil {
				fmt.Println("Unable to get prefix based on provided address")
				return
			}
		} else {
			ips, err := net.DefaultResolver.LookupIP(context.Background(), "ip6", "ipv4only.arpa.")
			if err != nil {
				fmt.Println("Unable to discover DNS64 prefix according to RFC7050")
				return
			}
			ip = ips[0]
		}
		ip[12] = 0
		ip[13] = 0
		ip[14] = 0
		ip[15] = 0
		subnet := 0
		for i := 15; i >= 0; i-- {
			if ip[i] != 0 {
				break
			}
			subnet += 8
		}
		fmt.Printf("%s/%d\n", ip.String(), subnet)
	},
}

func init() {
	RootCmd.AddCommand(&GetPrefix)
}
