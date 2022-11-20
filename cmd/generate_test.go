package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrint(t *testing.T) {
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	generateItem("")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = origStdout

	assert.Equal(t, "Generating ...\n", string(out))
}
