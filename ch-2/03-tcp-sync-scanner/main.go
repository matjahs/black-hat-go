package main

import (
	"fmt"
	"log"
	"net"
	"sort"
	"time"

	"github.com/pterm/pterm"
)

var (
	timeout = time.Second * 2
	dialer  = net.Dialer{Timeout: timeout}
)

type Result struct {
	port int
	open bool
}

func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := dialer.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		if err := conn.Close(); err != nil {
			log.Fatalln("failed to close connection")
		}
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	bar, err := pterm.DefaultProgressbar.WithTotal(1024).WithTitle("Scanning ports").Start()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		bar.Increment()
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}
