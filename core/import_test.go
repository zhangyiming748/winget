package core

import (
	"testing"
)

// go test -v -timeout 10m -run TestImport
func TestImport(t *testing.T) {
	Import("C:\\Users\\zen\\Github\\winget\\core\\export.json")
}
