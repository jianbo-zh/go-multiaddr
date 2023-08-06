//go:build !android
// +build !android

package manet

import (
	"fmt"
	"net"
)

type defaultInterface struct{}

func (defaultInterface) InterfaceAddrs() ([]net.Addr, error) {
	netlog.Info("other InterfaceAddrs")
	fmt.Println("other InterfaceAddrs")
	return net.InterfaceAddrs()
}

func getNetDriver() NetDriver {
	return &defaultInterface{}
}
