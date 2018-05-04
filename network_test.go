package tools

import (
	"testing"
)

func TestTcpLocalPortIsOpen(t *testing.T) {
	//It is very unlikely that port TCP/64000 is open
	if TcpLocalPortIsOpen(64000) {
		t.Error("Really? Port TCP/64000 open?")
	}
	//It is likely that port TCP/22 (SSH) is open
	if TcpLocalPortIsOpen(22) {
		t.Log("Port TCP/22 is open, as is likely the case, except on Termux/Android")
	}
}
