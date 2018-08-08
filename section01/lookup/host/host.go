package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	host string
)

func main() {
	flag.StringVar(&host, "host", "localhost", "hostname to resolve")
	flag.Parse()

	addrs, err := net.LookupHost(host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addrs)
}
