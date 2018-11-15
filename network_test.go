package tools

import (
	"fmt"
	"net"
	"testing"
)

func TestTcpLocalPortIsOpen(t *testing.T) {
	fmt.Println("[network]")
	l, err := net.Listen("tcp", ":61970")
	if err != nil {
		t.Error(err)
	}
	defer l.Close()
	if TcpLocalPortIsOpen(61970) {
		fmt.Println("> Port TCP/61970 is open")
	} else {
		t.Error("> Port TCP/61970 not open. There is problem")
	}
}
