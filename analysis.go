package goInspector

import (
	"encoding/json"
	"github.com/goretk/gore"
	"path/filepath"
)

type Vendor struct {
	Name string `json:"name"`
}

// InspectLibraries return the list with vendors
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

	std2, err := f.GetVendors()
	if err != nil {
		return nil, err
	}

	return std2, nil
}

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
