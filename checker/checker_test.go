package checker

import (
	"testing"
)

func BenchmarkConcurrentScanning(b *testing.B) {
	address := "scanme.nmap.org"
	network := "tcp"

	ConcurrentScanning(address, network)
}

func BenchmarkNonConcurrentScanning(b *testing.B) {
	address := "scanme.nmap.org"
	network := "tcp"

	NonConcurrentScanning(address, network)
}

func TestNonConcurrentScanning(t *testing.T) {
	address := "scanme.nmap.org"
	network := "tcp"

	t.Run("A1", func(t *testing.T) {
		address = "scanme.nmap.org"
		network = "tcp"
		ConcurrentScanning(address, network)
	})
	t.Run("A2", func(t *testing.T) {
		address = "ya.ru"
		network = "tcp"
		ConcurrentScanning(address, network)
	})
}
