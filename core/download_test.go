package core

import (
	"testing"
)

// go test -v -run TestDownload
func TestDownload(t *testing.T) {
	ids := []string{"GoLang.Go"}
	Download(ids, "C:\\Users\\zen\\Github\\winget\\core")
}
