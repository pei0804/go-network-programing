package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s network-type hostname\n", os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	name := os.Args[2]

	addr, err := net.ResolveTCPAddr(networkType, name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is ", addr.String())
	os.Exit(0)
}
