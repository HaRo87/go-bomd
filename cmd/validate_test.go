package cmd

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePrint(t *testing.T) {
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	validateItem("")

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = origStdout

	assert.Equal(t, "Validating ...\n", string(out))
}
