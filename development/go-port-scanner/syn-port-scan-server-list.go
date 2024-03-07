package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

const (
	protocolTCP = "tcp"
	timeout     = 2 * time.Second
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		printUsage()
		return
	}

	port := os.Args[1]
	filename := os.Args[2]

	ipList, err := readIPsFromFile(filename)
	if err != nil {
		fmt.Println("Error reading IP addresses:", err)
		return
	}

	var wg sync.WaitGroup

	for _, ip := range ipList {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			synScan(ip, port)
		}(ip)
	}

	wg.Wait()
}

// FUNCTIONS

func printUsage() {
	// Your printUsage function remains unchanged
}

func readIPsFromFile(filename string) ([]string, error) {
	// Your readIPsFromFile function remains unchanged
}

// SYN SCAN FUNCTIONS
func synScan(ip, port string) {
	target := fmt.Sprintf("%s:%s", ip, port)

	// Craft a TCP SYN packet
	packet := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{FixLengths: true, ComputeChecksums: true}
	err := gopacket.SerializeLayers(packet, opts,
		&layers.Ethernet{},
		&layers.IPv4{
			SrcIP:    net.IP{127, 0, 0, 1},
			DstIP:    net.ParseIP(ip),
			Protocol: layers.IPProtocolTCP,
		},
		&layers.TCP{
			SrcPort: layers.TCPPort(12345),
			DstPort: layers.TCPPort(atoi(port)),
			SYN:     true,
		},
	)
	if err != nil {
		fmt.Println("Error crafting SYN packet:", err)
		return
	}

	// Send the packet and wait for a response
	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		fmt.Println("Error listening for response:", err)
		return
	}
	defer conn.Close()

	_, err = conn.WriteTo(packet.Bytes(), &net.IPAddr{IP: net.ParseIP(ip)})
	if err != nil {
		fmt.Println("Error sending SYN packet:", err)
		return
	}

	buf := make([]byte, 1500)
	conn.SetReadDeadline(time.Now().Add(timeout))
	_, _, err = conn.ReadFrom(buf)
	if err != nil {
		fmt.Printf("Port %s/tcp is closed on %s\n", port, ip)
	} else {
		fmt.Printf("Port %s/tcp is open on %s\n", port, ip)
	}
}

func atoi(s string) uint16 {
	i, _ := strconv.Atoi(s)
	return uint16(i)
}
