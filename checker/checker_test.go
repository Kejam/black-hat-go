package main

import "testing"

func TestConcurrentScanning(t *testing.T) {
	address := "scanme.nmap.org"
	network := "tcp"

	ConcurrentScanning(address, network)
}

func TestNonConcurrentScanning(t *testing.T) {
	address := "scanme.nmap.org"
	network := "tcp"

	NonConcurrentScanning(address, network)
}
