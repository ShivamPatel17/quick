package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

// Define a struct to store flag values
type Options struct {
	Type string `short:"t" long:"type" description:"Specify the type"`
}

func main() {
	var opts Options

	// Parse command-line arguments into the struct
	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		os.Exit(1)
	}

	// Print the value of --type if provided
	fmt.Println("Type:", opts.Type)
}
