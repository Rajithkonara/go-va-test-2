package main

import (
	"fmt"
	"os"

	"github.com/arsham/rainbow/rainbow"
)

func main() {
	fmt.Println("Hello va scans")

	// ...
	l := rainbow.Light{
		Reader: os.Stdin,  // to read from
		Writer: os.Stdout, // to write to
	}
	l.Paint() // will rainbow everything it reads from reader to writer.
}
