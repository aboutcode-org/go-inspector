//
// Copyright (c) nexB Inc. and others. All rights reserved.
// ScanCode is a trademark of nexB Inc.
// SPDX-License-Identifier: Apache-2.0
// See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
// See https://github.com/nexB/go-inspector for support or download.
// See https://aboutcode.org for more information about nexB OSS projects.
//

package go_inspector

import (
	"debug/buildinfo"
	"encoding/json"
	"errors"
)

type Package struct {
	Path    string `json:"path"`
	Version string `json:"version"`
	Sum     string `json:"sum"`
}

// InspectLibrariesFunc interface to abstract the inspection of libraries
type InspectLibrariesFunc func(filepath string) ([]Package, error)

// InspectLibraries return all 3rd party packages used by the binary
func (i InspectLibrariesFunc) InspectLibraries(filepath string) ([]Package, error) {
	var packages []Package
	bi, err := buildinfo.ReadFile(filepath)
	if err != nil {
		return packages, errors.New("file not found")
	}
	for _, v := range bi.Deps {
		item := Package{Path: v.Path, Version: v.Version, Sum: v.Sum}
		packages = append(packages, item)
	}
	return packages, nil
}

// ConvertToJSONWithPosixPaths allows to inspect libraries at the specified path, convert the file paths to POSIX format, and obtain the JSON representation of the library information with POSIX paths.
func ConvertToJSONWithPosixPaths(path string, inspector InspectLibrariesFunc) ([]byte, error) {
	vendors, err := inspector.InspectLibraries(path)
	if err != nil {
		return nil, err
	}

	// Marshal the data to JSON
	dataJSON, err := json.Marshal(vendors)
	if err != nil {
		return nil, err
	}

	return dataJSON, nil
}
