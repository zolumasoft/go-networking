package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage program 1.1.1.1 where 1.1.1.1 is IP address")
		return
	}

	ip := net.ParseIP(os.Args[1])
	if ip != nil {
		fmt.Printf("%v OK\n", ip)
	} else {
		fmt.Println("bad address")
	}
}
