package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	cmd := os.Args[1]
	name := os.Args[2]

	var dns []string
	switch cmd {
	case "txt":
		dns, _ = net.LookupTXT(name)
	case "host":
		dns, _ = net.LookupHost(name)
	}
	fmt.Println(dns)
}
