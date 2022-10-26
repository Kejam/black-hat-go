package checker

import (
	"fmt"
	"log"
	"testing"
)

func TestCustomIOProxy(t *testing.T) {
	var reader CustomMockReader
	var writer CustomMockWriter

	input := make([]byte, 4096)

	read, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data", err)
	}
	fmt.Println("Read ", read, " bytes from stdin")

	write, err := writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data", err)
	}
	fmt.Println("Wrote ", write, " bytes from stdin")
}
