package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	cmd := os.Args[1]

	dns, _ := net.LookupTXT(cmd)
	fmt.Printf("%v", dns)
}
