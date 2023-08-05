//
// Copyright (c) nexB Inc. and others. All rights reserved.
// ScanCode is a trademark of nexB Inc.
// SPDX-License-Identifier: Apache-2.0
// See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
// See https://github.com/nexB/scancode-toolkit for support or download.
// See https://aboutcode.org for more information about nexB OSS projects.
//

package go_inspector

import (
	"encoding/json"
	"github.com/goretk/gore"
	"path/filepath"
)

type Vendor struct {
	Name string `json:"name"`
}

// InspectLibraries return all 3rd party packages used by the binary
func InspectLibraries(filepath string) ([]*gore.Package, error) {
	f, err := gore.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer func(f *gore.GoFile) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	packages, err := f.GetVendors()
	if err != nil {
		return nil, err
	}

	return packages, nil
}

// ConvertToJSONWithPosixPaths allows to inspect libraries at the specified path, convert the file paths to POSIX format, and obtain the JSON representation of the library information with POSIX paths.
func ConvertToJSONWithPosixPaths(path string) ([]byte, error) {
	vendors, err := InspectLibraries(path)
	if err != nil {
		return nil, err
	}

	// Convert the data to use POSIX paths
	for i := range vendors {
		vendors[i].Name = filepath.ToSlash(vendors[i].Name)
	}

	// Marshal the data to JSON
	dataJSON, err := json.Marshal(vendors)
	if err != nil {
		return nil, err
	}

	return dataJSON, nil
}
