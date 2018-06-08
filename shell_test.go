package tools

import (
	"fmt"
	"testing"
)

func TestBashExec(t *testing.T) {
	var err error
	commands := []string{
		"date",
		"date +%F",
		"date | wc -c",
		fmt.Sprintf("echo %q", "hello quoted"),
	}
	for _, command := range commands {
		if err = BashExec(command); err != nil {
			t.Errorf("command %q failed! %v\n", command, err)
		}
		t.Logf("successful execution of: %q\n", command)
	}
}
