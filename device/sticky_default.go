//go:build !linux

package device

import (
	"github.com/cawidtu/wireguard-go/conn"
	"github.com/cawidtu/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
