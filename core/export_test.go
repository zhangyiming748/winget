package core

import (
	"testing"
)

// go test -v -timeout 30m -run TestExport
func TestExport(t *testing.T) {
	Export("C:\\Users\\zen\\Github\\winget\\core")
}
