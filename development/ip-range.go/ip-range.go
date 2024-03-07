package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run ip-range.go <baseIP> <netmask>")
		os.Exit(1)
	}

	baseIP := os.Args[1]
	netmask := os.Args[2]

	ip, ipNet, err := net.ParseCIDR(baseIP + "/" + netmask)
	if err != nil {
		fmt.Println("Error parsing CIDR:", err)
		os.Exit(1)
	}

	fmt.Println("IP Range:")
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incrementIP(ip) {
		fmt.Println(ip)
	}
}

// incrementIP increments the IP address by 1
func incrementIP(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}
