package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	protocolTCP = "tcp"
	protocolUDP = "udp"
	timeout     = 2 * time.Second
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		printUsage()
	} else if len(os.Args) == 3 {
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
				scanPort(ip, port, protocolTCP)
				scanPort(ip, port, protocolUDP)
			}(ip)
		}

		wg.Wait()
	} else if len(os.Args) == 4 {
		port := os.Args[1]
		protocol := os.Args[2]
		filename := os.Args[3]

		if protocol != protocolTCP && protocol != protocolUDP {
			fmt.Println("Invalid protocol specified. Use 'tcp' or 'udp'.")
			return
		}

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
				scanPort(ip, port, protocol)
			}(ip)
		}

		wg.Wait()
	}
}

func printUsage() {
	fmt.Println("\nThis program will scan TCP and UDP ports out of a list of IP addresses from text/ASCII format file.\n")
	fmt.Println("By default, it will scan for both TCP and UDP if no protocol is specified.")
	fmt.Println("Optionally, you can specify either TCP or UDP protocol to limit your scan to that specific protocol.\n")
	fmt.Println("⚠️ WARNING ⚠️")
	fmt.Println("This program boldly engages in a lively TCP handshake dance, showcasing its impressive SYN, SYN-ACK, and ACK moves (a full TCP handshake scan).")
	fmt.Println("Beware! Such enthusiastic network exploration might prompt security devices, IDS, firewalls, and other guardians of the digital realm to raise an eyebrow or two. 🤨")
	fmt.Println("For a safer journey through the electronic dance floor, it's highly recommended to consult with your trusty security team. Avoiding curses and maintaining network harmony are their specialties! 🕵️‍♂️✨\n")

	fmt.Println("Usage:\ngo-port-scanner <port-number> <filename-with-list-of-ip-addresses>")
	fmt.Println("go-port-scanner <port-number> <tcp-or-udp> <filename-with-list-of-ip-addresses>\n")
	fmt.Println("Example:\ngo-port-scanner 22 servers.txt\n")
	fmt.Println("go-port-scanner 22 tcp servers.txt")
	fmt.Println("go-port-scanner 53 udp servers.txt\n")

	//fmt.Printf("Total arguments passed: %d\n\n", len(os.Args))
}

func readIPsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ips []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := strings.TrimSpace(scanner.Text())
		if ip != "" {
			ips = append(ips, ip)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ips, nil
}

func scanPort(ip, port, protocol string) {
	address := fmt.Sprintf("%s:%s", ip, port)
	conn, err := net.DialTimeout(protocol, address, timeout)
	if err == nil {
		fmt.Printf("Port %s/%s is open on %s\n", port, protocol, ip)
		conn.Close()
	} else {
		fmt.Printf("Port %s/%s is closed on %s\n", port, protocol, ip)
	}
}
