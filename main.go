package main

import (
	"fmt"
	"net"
	"time"
)

func scanPort(targetIP string, port int, timeout time.Duration) bool {
	target := fmt.Sprintf("%s:%d", targetIP, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		fmt.Printf("[ERROR] error scanning %s: %s", targetIP, err)
		return false
	}
	conn.Close()
	return true
}

func main() {
	targetIP := "192.168.1.1" // your target IP
	timeout := 2 * time.Second

	for port := 1; port <= 65535; port++ {
		if scanPort(targetIP, port, timeout) {
			fmt.Printf("[SUCCESS] port %d is open\n", port)
		}
	}
}
