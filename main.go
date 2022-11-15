package main

import (
	"github.com/tantalor93/synth6/cmd"
)

func main() {
	cmd.Execute()
}

//
// func main() {
// 	addr := net.ParseIP(os.Args[2])
// 	if addr == nil {
// 		fmt.Printf("provided address %s is not IP", os.Args[2])
// 		os.Exit(1)
// 	}
// 	addr = addr.To4()
// 	if addr == nil {
// 		fmt.Printf("provided address %s is not IP", os.Args[2])
// 		os.Exit(1)
// 	}
//
// 	_, prefix, err := net.ParseCIDR(os.Args[1])
// 	if err != nil {
// 		fmt.Printf("failed to parse net %s", os.Args[1])
// 		os.Exit(1)
// 	}
//
// 	n, _ := prefix.Mask.Size()
// 	// Assumes prefix has been validated during setup
// 	v6 := make([]byte, 16)
// 	i, j := 0, 0
//
// 	for ; i < n/8; i++ {
// 		v6[i] = prefix.IP[i]
// 	}
// 	for ; i < 8; i, j = i+1, j+1 {
// 		v6[i] = addr[j]
// 	}
// 	if i == 8 {
// 		i++
// 	}
// 	for ; j < 4; i, j = i+1, j+1 {
// 		v6[i] = addr[j]
// 	}
//
// 	fmt.Println(net.IPNet{IP: v6}.IP)
// }
