package findbin

import (
	"os"
	"path/filepath"
	"testing"
)

func FindBin(t *testing.T) string {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	binary := filepath.Base(dir)

	absBinary := filepath.Join(dir, binary)

	_, err = os.Open(absBinary)
	if err != nil {
		t.Fatal("Looks like there is no " + binary + " binary here. Did you run go build?")
	}

	t.Log("Found a binary named " + binary)

	return absBinary
}
