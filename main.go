package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, addr := range addrs {
		// Check the address type and skip the loopback address
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println("Your IP address is:", ipNet.IP.String())
				return
			}
		}
	}

	fmt.Println("No IP address found")
}
