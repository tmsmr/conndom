package wifi

import (
	"errors"

	gonm "github.com/Wifx/gonetworkmanager/v3"
	"github.com/godbus/dbus/v5"
)

var (
	ErrInvalidStaDevice = errors.New("invalid STA device")
)

type StaDevice struct {
	Path       dbus.ObjectPath
	Name       string
	MacAddress string
	Driver     string
}

func (m *Manager) GetStaDevice() (*StaDevice, error) {
	objects, err := m.nm.GetDevices()
	if err != nil {
		return nil, err
	}
	for _, object := range objects {
		name, err := object.GetPropertyInterface()
		if err != nil {
			return nil, err
		}
		if name != m.conf.StaWifiInterface {
			continue
		}
		deviceType, err := object.GetPropertyDeviceType()
		if err != nil {
			return nil, err
		}
		if deviceType != gonm.NmDeviceTypeWifi {
			return nil, ErrInvalidStaDevice
		}
		driver, err := object.GetPropertyDriver()
		if err != nil {
			return nil, err
		}
		macAddr, err := object.GetPropertyHwAddress()
		if err != nil {
			return nil, err
		}
		return &StaDevice{
			Path:       object.GetPath(),
			Name:       name,
			MacAddress: macAddr,
			Driver:     driver,
		}, nil
	}
	return nil, ErrInvalidStaDevice
}
