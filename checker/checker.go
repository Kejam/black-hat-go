package checker

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
)

func ConcurrentScanning(address string, networkType string) {
	ports := make(chan int, 10)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go dialSync(ports, results, address, networkType)
	}

	go func() {
		for i := 0; i < 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, openport := range openports {
		fmt.Println("Port ", openport, " is open")
	}
}

func dialSync(ports, result chan int, address, networkType string) {
	for port := range ports {
		builder := strings.Builder{}
		builder.WriteString(address)
		builder.WriteString(":")
		builder.WriteString(strconv.Itoa(port))
		addressWithPort := builder.String()
		conn, err := net.Dial(networkType, addressWithPort)
		if err != nil {
			//fmt.Println("Error dial to: ", addressWithPort)
			result <- 0
			continue
		} else {
			//fmt.Println("Success dial to: ", addressWithPort)
			result <- port
		}
		err = conn.Close()
		if err != nil {
			fmt.Println("Error closed ", err)
			continue
		}
	}
}

func NonConcurrentScanning(address string, networkType string) {
	for i := 0; i < 1024; i++ {
		builder := strings.Builder{}
		builder.WriteString(address)
		builder.WriteString(":")
		builder.WriteString(strconv.Itoa(i))
		addressWithPort := builder.String()
		dial(addressWithPort, networkType)
	}
}

func dial(address string, networkType string) {
	conn, err := net.Dial(networkType, address)
	if err != nil {
		//fmt.Println("Error dial to: ", address)
		return
	} else {
		fmt.Println("Success dial to: ", address)
	}
	err = conn.Close()
	if err != nil {
		// Skip error
		return
	}
}
