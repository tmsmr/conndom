package wifi

import (
	"time"

	gonm "github.com/Wifx/gonetworkmanager/v3"
)

const (
	apScanTimeoutSeconds = 30
)

type NetworkFrequency struct {
	Frequency  uint32
	Channel    uint8
	Strength   uint8
	MaxBitrate uint32
}

type Network struct {
	Private     bool
	Frequencies []NetworkFrequency
}

func (d StaDevice) Scan() (map[string]Network, error) {
	wd, err := gonm.NewDeviceWireless(d.Path)
	if err != nil {
		return nil, err
	}
	last, err := wd.GetPropertyLastScan()
	if err != nil {
		return nil, err
	}
	err = wd.RequestScan()
	if err != nil {
		return nil, err
	}
	for i := 0; i < apScanTimeoutSeconds*10; i++ {
		current, err := wd.GetPropertyLastScan()
		if err != nil {
			return nil, err
		}
		if current > last {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	aps, err := wd.GetAccessPoints()
	if err != nil {
		return nil, err
	}
	networks := make(map[string]Network)
	for _, ap := range aps {
		ssid, err := ap.GetPropertySSID()
		if err != nil {
			return nil, err
		}
		frequency, err := ap.GetPropertyFrequency()
		if err != nil {
			return nil, err
		}
		strength, err := ap.GetPropertyStrength()
		if err != nil {
			return nil, err
		}
		maxBitrate, err := ap.GetPropertyMaxBitrate()
		if err != nil {
			return nil, err
		}
		flags, err := ap.GetPropertyFlags()
		if err != nil {
			return nil, err
		}
		network, ok := networks[ssid]
		if !ok {
			network = Network{Private: flags != 0}
		}
		network.Frequencies = append(network.Frequencies, NetworkFrequency{
			Frequency:  frequency,
			Strength:   strength,
			MaxBitrate: maxBitrate,
		})
		networks[ssid] = network
	}
	return networks, nil
}
