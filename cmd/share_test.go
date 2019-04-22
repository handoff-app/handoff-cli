package cmd

import "testing"

func TestErrorThrownWhenCommandCalledWithoutFilePath(t *testing.T) {
	var args []string
	if err := shareCmd.Args(shareCmd, args); err == nil {
		t.Error("share command should fail when no path passed")
	}
}


