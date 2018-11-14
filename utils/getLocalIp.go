package utils

import (
	"net"

	"github.com/pkg/errors"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, errors.New("can't get ip")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}
