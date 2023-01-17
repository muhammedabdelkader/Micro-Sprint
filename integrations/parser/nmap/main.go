package main

import (
	"fmt"
	"nmap"
)

func main() {
	nmapRun, err := nmap.ParseFile("sample.xml")
	if err != nil {
		fmt.Printf("Error parsing file: %v", err)
		return

	}

	fmt.Println("Scanner:", nmapRun.Scanner)
	fmt.Println("Arguments:", nmapRun.Args)

	for _, host := range nmapRun.Hosts {
		fmt.Println("Host:", host.Addresses[0].Addr)
		fmt.Println("Status:", host.Status.State)
		fmt.Println("Open Ports:")

		for _, port := range host.Ports {
			if port.State.State == "open" {
				fmt.Printf("  %s/%s\n", port.PortId, port.Protocol)

			}

		}

	}

}
