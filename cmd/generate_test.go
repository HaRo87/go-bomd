package cmd

import (
	"os"
	"path"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTemplateCmdInvalidFileReturnsError(t *testing.T) {
	hook := test.NewGlobal()
	err := executeCommand(rootCmd, "generate", "template", "-vvv")
	assert.Error(t, err)
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "😱 something went wrong", hook.LastEntry().Message)
	hook.Reset()
}

func TestGenerateTemplateCmdSuccess(t *testing.T) {
	hook := test.NewGlobal()
	tempDir := os.TempDir()
	filePath := path.Join(tempDir, "not_a_template.tmpl")
	err := executeCommand(rootCmd, "generate", "template", "-vvv", "-f", filePath)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "😎 everything seems to be fine", hook.LastEntry().Message)
	hook.Reset()
}
