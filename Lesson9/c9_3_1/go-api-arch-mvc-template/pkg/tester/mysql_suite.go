package tester

import (
	"fmt"
	"net"
	"time"
)

func CheckPort(host string, port int) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if conn != nil {
		conn.Close()
		return false
	}
	if err != nil {
		return true
	}
	return false
}

func WaitForPort(host string, port int, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if CheckPort(host, port) {
			return true
		}
		time.Sleep(1 * time.Second)
	}
	return false
}
