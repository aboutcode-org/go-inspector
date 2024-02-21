// Copyright (c) nexB Inc. and others. All rights reserved.
// ScanCode is a trademark of nexB Inc.
// SPDX-License-Identifier: Apache-2.0
// See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
// See https://github.com/nexB/go-inspector for support or download.
// See https://aboutcode.org for more information about nexB OSS projects.
package go_inspector

import (
	"testing"
)

func TestInspectLibraries(t *testing.T) {
	// Test case 1: Provide a valid filepath
	filepath := "examples/files/basic_golang_app/basicGoApp.exe"
	mockInspect := new(InspectLibrariesFunc)
	packages, err := mockInspect.InspectLibraries(filepath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check the length of the packages
	expectedNumPackages := 14
	if len(packages) != expectedNumPackages {
		t.Errorf("Expected %d packages, but got %d", expectedNumPackages, len(packages))
	}

	// Test case 2: Provide an invalid filepath
	invalidFilepath := "nonexistent/file"
	_, err = mockInspect.InspectLibraries(invalidFilepath)
	if err == nil {
		t.Errorf("Expected an error for an invalid filepath, but got none")
	}
}
func TestConvertToJSONWithPosixPaths(t *testing.T) {
	mockInspect := new(InspectLibrariesFunc)
	dataJSON, err := ConvertToJSONWithPosixPaths("examples/files/basic_golang_app/basicGoApp.exe", mockInspect.InspectLibraries)
	if err != nil {
		t.Fatalf("ConvertToJSONWithPosixPaths failed: %v", err)
	}

	// Check if dataJSON contains the expected JSON representation of the packages
	expectedJSON := `[{"path":"github.com/gin-contrib/sse","version":"v0.1.0","sum":"h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE="},{"path":"github.com/gin-gonic/gin","version":"v1.8.1","sum":"h1:4+fr/el88TOO3ewCmQr8cx/CtZ/umlIRIs5M4NTNjf8="},{"path":"github.com/go-playground/locales","version":"v0.14.0","sum":"h1:u50s323jtVGugKlcYeyzC0etD1HifMjqmJqb8WugfUU="},{"path":"github.com/go-playground/universal-translator","version":"v0.18.0","sum":"h1:82dyy6p4OuJq4/CByFNOn/jYrnRPArHwAcmLoJZxyho="},{"path":"github.com/go-playground/validator/v10","version":"v10.11.0","sum":"h1:0W+xRM511GY47Yy3bZUbJVitCNg2BOGlCyvTqsp/xIw="},{"path":"github.com/leodido/go-urn","version":"v1.2.1","sum":"h1:BqpAaACuzVSgi/VLzGZIobT2z4v53pjosyNd9Yv6n/w="},{"path":"github.com/mattn/go-isatty","version":"v0.0.14","sum":"h1:yVuAays6BHfxijgZPzw+3Zlu5yQgKGP2/hcQbHb7S9Y="},{"path":"github.com/pelletier/go-toml/v2","version":"v2.0.2","sum":"h1:+jQXlF3scKIcSEKkdHzXhCTDLPFi5r1wnK6yPS+49Gw="},{"path":"github.com/ugorji/go/codec","version":"v1.2.7","sum":"h1:YPXUKf7fYbp/y8xloBqZOw2qaVggbfwMlI8WM3wZUJ0="},{"path":"golang.org/x/crypto","version":"v0.0.0-20220722155217-630584e8d5aa","sum":"h1:zuSxTR4o9y82ebqCUJYNGJbGPo6sKVl54f/TVDObg1c="},{"path":"golang.org/x/net","version":"v0.0.0-20220725212005-46097bf591d3","sum":"h1:2yWTtPWWRcISTw3/o+s/Y4UOMnQL71DWyToOANFusCg="},{"path":"golang.org/x/text","version":"v0.3.7","sum":"h1:olpwvP2KacW1ZWvsR7uQhoyTYvKAupfQrRGBFM352Gk="},{"path":"google.golang.org/protobuf","version":"v1.28.0","sum":"h1:w43yiav+6bVFTBQFZX0r7ipe9JQ1QsbMgHwbBziscLw="},{"path":"gopkg.in/yaml.v2","version":"v2.4.0","sum":"h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY="}]`
	if string(dataJSON) != expectedJSON {
		t.Errorf("Expected JSON: %s, but got %s", expectedJSON, string(dataJSON))
	}
}
