// Print a DNS zone file information and commands as an aid to locate config file, options, and parameters related to a DNS zone

package main

import (
	"fmt"
	"os"
)

func main() {
	var zoneName, zonesDir, prefix, postfix string
	var filePath string
	var zoneView []string

	zonesDir = "/var/named/zones/"
	zoneView = []string{"internal", "external"}
	prefix = "db"
	postfix = "zone"

	// Check if at least two command-line arguments are provided (zone-name view-name)
	if len(os.Args) < 3 {
		printUsage()
		return
	}

	zoneName = os.Args[1]
	zoneView = []string{os.Args[2]}
	filePath = zonesDir + zoneView[0] + "/" + prefix + "." + zoneName + "." + postfix

	switch {
	case len(os.Args) == 3:
		printZoneInfo(zoneName, zonesDir, zoneView[0], prefix, postfix, filePath)

	case len(os.Args) == 4 && os.Args[3] == "config":
		printGrepCommand(zoneName, zoneView[0])

	case len(os.Args) == 4 && os.Args[3] == "commands":
		printZoneInfo(zoneName, zonesDir, zoneView[0], prefix, postfix, filePath+"\n")
		fmt.Println("Working with the following domain/zone: " + zoneName + "\n")
		printBindCommands(zoneName, zoneView[0], filePath)

	default:
		fmt.Print("Invalid argument or invalid number of arguments.\n")
	}
}

func printUsage() {
	fmt.Print("Please provide the 'zone name' and the 'view name' as command-line arguments.\n\n")
	fmt.Println("Example of valid arguments:\n")
	fmt.Println("<domain-name> <view-name>\t\t Display information about the location of the provided DNS zone")
	fmt.Println("<domain-name> <view-name> config\t Display suggested command to look for the DNS zone configuration")
	fmt.Println("<domain-name> <view-name> commands\t Display all the suggested commands to work with the provided DNS zone\n")
	fmt.Println("Ex1: example.net internal")
	fmt.Println("Ex2: example.net external")
	fmt.Println("Ex3: example.net internal config")
	fmt.Println("Ex4: example.net internal commands\n")
}

func printZoneInfo(zoneName, zonesDir, zoneView, prefix, postfix, filePath string) {
	result := fmt.Sprintf("Zone Name: %s\nZones Dir: %s\nView: %v\nFilename: %s.%s.%s\nAbsolute Path: %s",
		zoneName, zonesDir, zoneView, prefix, zoneName, postfix, filePath)
	fmt.Println(result)
}

func printGrepCommand(zoneName, zoneView string) {
	fmt.Println("Check the options and the configuration set for the zone")
	fmt.Printf("grep -Hnr '%s' /etc/named/primary/zones.active.%s* -A4\n", zoneName, zoneView)
}

func printBindCommands(zoneName, zoneView, filePath string) {
	fmt.Println("# Check the main BIND configuration. No return means everything is good.")
	fmt.Println("sudo named-checkconf /etc/named/primary/named.active.conf\n")
	fmt.Println("# Check the options and the configuration set for the zone")
	fmt.Printf("grep -Hnr '%s' /etc/named/primary/zones.active.%s* -A4\n\n", zoneName, zoneView)
	fmt.Printf("# Check if the zone file exists\nsudo ls -la %s\n\n", filePath)
	fmt.Printf("# Check the zone config syntax. No return means everything is good.\nsudo named-checkconf %s\n\n", filePath)
	fmt.Printf("# Freeze the zone\nsudo rndc freeze %s IN %s\n\n", zoneName, zoneView)
	fmt.Printf("# Edit the zone\nsudo vi %s\n\n", filePath)
	fmt.Printf("# Delete journal of the zone\nsudo rm -i %s.jnl\n\n", filePath)
	fmt.Printf("# Check the syntax of the zone\nsudo named-checkzone %s %s\n\n", zoneName, filePath)
	fmt.Printf("# Thaw the zone\nsudo rndc thaw %s IN %s\n\n", zoneName, zoneView)
	fmt.Printf("# Unfreeze zone\nsudo rndc unfreeze %s IN %s\n\n", zoneName, zoneView)
	fmt.Printf("# Reload the zone. Will not work on dynamic zone but won't harm either\nsudo rndc reload %s\n\n", filePath)
}
