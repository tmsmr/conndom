package network

import gonm "github.com/Wifx/gonetworkmanager/v3"

type WifiDevice struct {
	Name       string
	MacAddress string
	Driver     string
}

func (m *Manager) ListWifiDevices() ([]WifiDevice, error) {
	objects, err := m.nm.GetPropertyAllDevices()
	if err != nil {
		return nil, err
	}
	var devices []WifiDevice
	for _, object := range objects {
		deviceType, err := object.GetPropertyDeviceType()
		if err != nil {
			return nil, err
		}
		if deviceType != gonm.NmDeviceTypeWifi {
			continue
		}
		realDevice, err := object.GetPropertyReal()
		if err != nil {
			return nil, err
		}
		if !realDevice {
			continue
		}
		driver, err := object.GetPropertyDriver()
		if err != nil {
			return nil, err
		}
		macAddr, err := object.GetPropertyHwAddress()
		if err != nil {
			return nil, err
		}
		name, err := object.GetPropertyInterface()
		if err != nil {
			return nil, err
		}
		devices = append(devices, WifiDevice{
			Name:       name,
			MacAddress: macAddr,
			Driver:     driver,
		})
	}
	return devices, nil
}
