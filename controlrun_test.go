package controlrun

import (
	"testing"
)

// Test Access
func TestAccess(t *testing.T) {
	hash := Set("project-name", "appName/da39a3ee5e6b4b0d3255bfef95601890afd80709")
	println("Hash verification generated:", hash)

	err := Run(hash)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("OK")
	}
}
