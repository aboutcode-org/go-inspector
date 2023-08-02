package main

import (
	"fmt"
	goInspector "github.com/nexB/go-inspector"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %v \n", args[0])
		fmt.Printf("  Load executable file file, and\n")
		fmt.Printf("  print portions of its creation info data.\n")
		return
	}
	jsonData, err := goInspector.ConvertToJSONWithPosixPaths(args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
