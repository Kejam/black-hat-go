package checker

import (
	"fmt"
	"os"
)

type CustomReader struct {
}

type CustomMockReader struct {
}

type CustomWriter struct {
}

type CustomMockWriter struct {
}

func (сustomReader *CustomReader) Read(b []byte) (int, error) {
	fmt.Println("in >")
	return os.Stdin.Read(b)
}

func (сustomReader *CustomMockReader) Read(b []byte) (int, error) {
	fmt.Println("in >")
	return 15, nil
}

func (сustomWriter *CustomWriter) Write(b []byte) (int, error) {
	fmt.Println("out >")
	return os.Stdin.Write(b)
}

func (сustomWriter *CustomMockWriter) Write(b []byte) (int, error) {
	fmt.Println("out >")
	return 4096, nil
}
