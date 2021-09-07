package main

import (
	"fmt"
	"net"
)

func main() {

	dns, _ := net.LookupTXT("api.vimmo.app")
	fmt.Printf("%v", dns)
}
