package main

import (
	"fmt"
	"net"
	"os"
)

func txtlookup(name string) (dns []string) {
	dns, err := net.LookupTXT(name)

	if err != nil {
		panic(err)
	}
	return dns
}
func LookupHost(name string) (dns []string) {
	dns, err := net.LookupHost(name)

	if err != nil {
		panic(err)
	}
	return dns
}
func main() {
	cmd := os.Args[1]
	name := os.Args[2]

	var dns []string
	switch cmd {
	case "txt":
		dns = txtlookup(name)

	case "host":
		dns = LookupHost(name)
	}
	fmt.Println(name, dns)
}
