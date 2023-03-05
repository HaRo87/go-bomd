package cmd

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestValidateBomCmdInvalidFileReturnsError(t *testing.T) {
	hook := test.NewGlobal()
	err := executeCommand(rootCmd, "validate", "bom", "-vvv", "-f", "bom.json")
	assert.Error(t, err)
	assert.Equal(t, 2, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "😱 something went wrong", hook.LastEntry().Message)
	hook.Reset()
}

func TestValidateBomCmdInvalidLicensesReturnsError(t *testing.T) {
	hook := test.NewGlobal()
	err := executeCommand(rootCmd, "validate", "bom", "-vvv", "--licenseCheck", "-f", "../examples/boms/go-bomd-bom.json")
	assert.Error(t, err)
	assert.Equal(t, 4, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "😱 something went wrong", hook.LastEntry().Message)
	hook.Reset()
}

func TestValidateBomCmdNoLicenseCheckSuccess(t *testing.T) {
	hook := test.NewGlobal()
	// Weird workaround due to the fact that flags do not seem to
	// be automatically reset during test execution. More details
	// can be found here: https://github.com/spf13/cobra/issues/1419
	assert.NoError(t, validateBomCmd.Flags().Lookup("licenseCheck").Value.Set("false"))
	err := executeCommand(rootCmd, "validate", "bom", "-vvv", "-f", "../examples/boms/go-bomd-bom.json")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "😎 everything seems to be fine", hook.LastEntry().Message)
	hook.Reset()

}

func TestValidateBomCmdNoLicenseCheckButListMissingLicensesSuccess(t *testing.T) {
	hook := test.NewGlobal()
	// Weird workaround due to the fact that flags do not seem to
	// be automatically reset during test execution. More details
	// can be found here: https://github.com/spf13/cobra/issues/1419
	assert.NoError(t, validateBomCmd.Flags().Lookup("licenseCheck").Value.Set("false"))
	err := executeCommand(rootCmd, "validate", "bom", "-vvv", "--listMissing", "-f", "../examples/boms/go-bomd-bom.json")
	assert.NoError(t, err)
	assert.Equal(t, 14, len(hook.Entries))
	assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
	assert.Equal(t, "🤔 component: gopkg.in/yaml.v3 is missing license information", hook.LastEntry().Message)
	hook.Reset()

}
