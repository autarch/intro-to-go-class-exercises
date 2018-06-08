package findbin

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func FindBin(t *testing.T) string {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	binary := filepath.Base(dir)
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	absBinary := filepath.Join(dir, binary)

	_, err = os.Open(absBinary)
	if err != nil {
		t.Fatalf("Looks like there is no %s binary here. Did you run go build?", binary)
	}

	t.Logf("Found a binary named %s", binary)

	return absBinary
}
