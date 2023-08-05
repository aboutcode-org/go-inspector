//
// Copyright (c) nexB Inc. and others. All rights reserved.
// ScanCode is a trademark of nexB Inc.
// SPDX-License-Identifier: Apache-2.0
// See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
// See https://github.com/nexB/scancode-toolkit for support or download.
// See https://aboutcode.org for more information about nexB OSS projects.
//

package main

import (
	"fmt"
	goinspector "github.com/nexB/go-inspector"
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
	jsonData, err := goinspector.ConvertToJSONWithPosixPaths(args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
