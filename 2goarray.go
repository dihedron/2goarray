// Simple utility to convert a file into a Go byte array

// Clint Caywood

// http://github.com/cratonica/2goarray
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	packageName := flag.String("package", "", "the package to use for the file (default: main)")
	arrayName := flag.String("array", "Data", "the name of the array to create (default: Data)")
	helpFlag := flag.Bool("help", false, "print this help message and quit")
	versionFlag := flag.Bool("version", false, "print version information and quit")

	if *versionFlag {
		fmt.Println("2goarray v1.0.0")
		os.Exit(0)
	}

	if *helpFlag {
		fmt.Printf(`
2goarray v1.0.0  - a utility to turn any resource into an embeddable Golang array

usage: 
      2goarray [-help] [-version] [-array <array name>] [-package <package_name>]")
where
      -help     print this help message and quit
      -version  print version information and quit
      -array    the name of the array to generate (default: Data)
      -package  the optional name of the Golang package  (default: none)

`)
		os.Exit(0)
	}

	if isTerminal() {
		fmt.Printf("\nPlease pipe the file you wish to encode into stdin\n\n")
		os.Exit(0)
	}

	flag.Parse()

	if *packageName != "" {
		fmt.Printf("package %s\n\n", *packageName)
	}

	fmt.Printf("var %s []byte = []byte{", *arrayName)
	buf := make([]byte, 1)
	var err error
	var totalBytes uint64
	var n int
	for n, err = os.Stdin.Read(buf); n > 0 && err == nil; {
		if totalBytes%12 == 0 {
			fmt.Printf("\n\t")
		}
		fmt.Printf("0x%02x, ", buf[0])
		totalBytes++
		n, err = os.Stdin.Read(buf)
	}
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
	fmt.Print("\n}\n")
}
