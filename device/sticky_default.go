//go:build !linux

package device

import (
	"github.com/cawidtu/notwireguard-go/conn"
	"github.com/cawidtu/notwireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
