package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type portscan struct {
	ip string
}

func runscan(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return
	}
	conn.Close()
	fmt.Println(port, "is open")
}

func (scan *portscan) run(min, max int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := min; port <= max; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			runscan(scan.ip, port, timeout)
		}(port)
	}
}

func main() { // get the goodies
	var host string
	var startport int
	var endport int

	fmt.Println("\nHost To Scan: ")
	fmt.Scanln(&host)
	fmt.Println("\nPort To Start Scanning At: ")
	fmt.Scanln(&startport)
	fmt.Println("\nPort To End The Scan At: ")
	fmt.Scanln(&endport)

	if startport < 0 || endport >= 65536 {
		fmt.Println("\nPort Range Invalid, Ports Must Be Between 0 And 65535")
		return
	}

	fmt.Println("\nScan Started\n======================\n")

	scan := &portscan{
		ip: host,
	}
	scan.run(startport, endport, 500*time.Millisecond)
}
