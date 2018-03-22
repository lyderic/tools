package tools

import (
	"net"
)

func TcpLocalPortIsOpen(int port) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return false
	} else {
		defer conn.Close()
		return true
	}
}
