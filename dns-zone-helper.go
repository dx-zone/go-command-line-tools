// Print a DNS zone file information and commands as an aid to locate config file, options, and parameters related to a DNS zone

package main

import (
	"fmt"
	"os"
	"time"
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
	fmt.Println("# Best practice: Create a backup of the BIND configuration file and the specific zone file, storing it with a timestamp.")
	timestamp := time.Now().Format("20060102-150405")
    fmt.Println("sudo tar -czvf /srv/dns-zone-" + zoneName + "-backup-" + timestamp + ".tar.gz /etc/named.conf /var/named/zones/" + zoneView + "/db." + zoneName + ".zone\n")
	fmt.Println("# Copy the BIND configuration and zone files to a temporary directory for safety.")
	fmt.Printf("sudo cp -a /etc/named.conf /var/named/zones/%s/db.%s.zone /tmp/\n\n", zoneView, zoneName)
	fmt.Println("# Validate the main BIND configuration file. If no errors are output, the configuration is correct.")
	fmt.Println("sudo named-checkconf /etc/named/primary/named.active.conf\n")
	fmt.Println(`# Search for '` + zoneName + `' within the active zone configurations and display 4 lines after each match to provide context. This is useful for validating the zone type configuration and ensuring correct settings.`)
	fmt.Printf("sudo grep -Hnr 'zone \"%s' /etc/named/primary/zones.active.*%s* -A4\n\n", zoneName, zoneView)
    fmt.Println("\nNOTE: Proceed with the following commands if the zone type is master (slave zones can't be edited as the records will be overwritten by the master DNS for the zone).\n")
	fmt.Printf("# List details for the specified zone file to confirm its existence and permissions.\nsudo ls -la %s\n\n", filePath)
	fmt.Println("# Check the syntax and validate the that the main BIND configuration is free of errors. No output indicates no errors.\nsudo named-checkconf \\etc\\named.conf\n")
	fmt.Printf("# Freeze updates to the zone '%s' to safely edit the zone file.\nsudo rndc freeze %s IN %s\n\n", zoneName, zoneName, zoneView)
	fmt.Printf("# Open the zone file in the Vi editor for manual modifications. Save the changes by pressing ':wq' and then Enter.\nsudo vi %s\n\n", filePath)
	fmt.Printf("# Safely remove the journal file associated with the zone to reset DNS record state tracking.\nsudo rm -i %s.jnl\n\n", filePath)
	fmt.Printf("# Check syntax and validate that the edited zone file is free of errors, ensuring all records are correctly formatted.\nsudo named-checkzone %s %s\n\n", zoneName, filePath)
	fmt.Printf("# Thaw the zone to resume automatic processing of updates.\nsudo rndc thaw %s IN %s\n\n", zoneName, zoneView)
	fmt.Printf("# Deprecated command: 'unfreeze' is no longer used, replaced by 'thaw'.\nsudo rndc unfreeze %s IN %s\n\n", zoneName, zoneView)
	fmt.Printf("Instruct BIND to reload its configuration and zone files without the need to completely restart the server service.\nsudo rndc reload %s IN %s %s\n\n",zoneName, zoneView, filePath)
    fmt.Println("\nNOTE: Proceed with the following command if the zone type is slave.\n")
	fmt.Printf("On a Slave BIND DNS - instructs the BIND DNS slave server to refresh the DNS zone within the view, updating its data from the master server\nsudo rndc refresh %s IN %s\n\n",zoneName, zoneView)
}
