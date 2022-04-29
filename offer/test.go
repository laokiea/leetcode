package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "h6.irisgw.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conn)
}
