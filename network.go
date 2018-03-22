package tools

import (
  "fmt"
	"net"
)

func TcpLocalPortIsOpen(port int) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return false
	} else {
		defer conn.Close()
		return true
	}
}
