package tools

import (
	"fmt"
	"net"
	"testing"
)

func TestTcpLocalPortIsOpen(t *testing.T) {
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

func TestTcpPortIsOpen(t *testing.T) {
	host := "google.com"
	port := 443
	if TcpPortIsOpen(host, port) {
		fmt.Printf("Port %d on %q is in use\n", port, host)
	} else {
		t.Errorf("> Is %s down?", host)
	}
}
