package controlrun

import (
	"testing"
)

// Test Access
func TestAccess(t *testing.T) {
	Set("project-name", "appName/a9zaSyB7sRpNWuVbNb")

	access, err := Run()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(access)
	}
}
