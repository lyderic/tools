package tools

import (
	"fmt"
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

func TestBashExecOutput(t *testing.T) {
	stdout, stderr, err := BashExecCapture(scriptOutput, "hello output")
	if err != nil {
		t.Errorf("script %q failed! %v\n", script, err)
	}
	fmt.Printf("OUT [%d]> %q\n", len(stdout), string(stdout))
	fmt.Printf("ERR [%d]> %q\n", len(stderr), string(stderr))
}

/*
func TestLongOutput(t *testing.T) {
	var err error
	if err = BashExec("ls -R /"); err != nil {
		t.Errorf("script %q failed! %v\n", script, err)
	}
}
*/

var script = `#!/bin/bash

echo -n "Date: "
date +%F
echo "Input: ${1}"
echo "@=${@}"
`

var scriptOutput = `#!/bin/bash

echo $@
echo $@ >&2
`
