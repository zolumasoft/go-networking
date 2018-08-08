package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	addr string
)

func main() {
	flag.StringVar(&addr, "addr", "127.0.0.1", "host address to lookup")
	flag.Parse()

	names, err := net.LookupAddr(addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(names)
}
