package tools

import (
	"testing"
)

func TestBashExec(t *testing.T) {
	var err error
	if err = BashExec(script); err != nil {
		t.Errorf("script %q failed! %v\n", script, err)
	}
	if err = BashExec(script, "foo", "bar", "baz"); err != nil {
		t.Errorf("script %q failed! %v\n", script, err)
	}
}

var script = `#!/bin/bash

echo -n "Date: "
date +%F
echo "Input: ${1}"
echo "@=${@}"
`
