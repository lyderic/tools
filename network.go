package tools

import (
	"fmt"
	"net"
)

/* probe TCP port. true if port is in use on host */
func TcpPortIsOpen(host string, port int) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return false
	} else {
		defer conn.Close()
		return true
	}
}

/* true if TCP localhost:<port> is available */
func TcpLocalPortIsOpen(port int) bool {
	return TcpPortIsOpen("localhost", port)
}
