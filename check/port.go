package check

import (
	"fmt"
	"net"
	"time"

	"enkodr.dev/status/config"
)

// Port identifies the port structure to check
type Port struct{}

// Check will test if the port is open
func (p Port) Check(config config.Check) (bool, error) {
	target := fmt.Sprintf("%s:%d", config.Address, config.Port)
	timeout := time.Second
	conn, err := net.DialTimeout(config.Protocol, target, timeout)
	if err != nil {
		return false, err
	}
	conn.Close()
	return true, nil
}
