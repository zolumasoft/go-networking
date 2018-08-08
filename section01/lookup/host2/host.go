package main

import (
	"context"
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

	res := net.Resolver{PreferGo: true}
	addrs, err := res.LookupHost(context.Background(), host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addrs)
}
