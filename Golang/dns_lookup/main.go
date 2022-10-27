package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ips, err := net.LookupIP("sendgrid.net")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	for _, ip := range ips {
		fmt.Printf("sendgrid.net. IN A %s\n", ip.String())
	}
}
