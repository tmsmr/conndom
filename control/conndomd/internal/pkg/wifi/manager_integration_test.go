package wifi_test

import (
	"testing"

	"github.com/tmsmr/conndom/conndomd/internal/pkg/env"
	"github.com/tmsmr/conndom/conndomd/internal/pkg/wifi"
)

const forwardedDbusSocket = "/tmp/conndom_dbus_system.sock"

func TestManager_GetStaDevice(t *testing.T) {
	m, err := wifi.NewManager(env.Spec{
		DbusSystemBusAddress: forwardedDbusSocket,
		StaWifiInterface:     "wlan1",
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.GetStaDevice()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDevice_Scan(t *testing.T) {
	m, err := wifi.NewManager(env.Spec{
		DbusSystemBusAddress: forwardedDbusSocket,
		StaWifiInterface:     "wlan1",
	})
	if err != nil {
		t.Fatal(err)
	}
	device, err := m.GetStaDevice()
	if err != nil {
		t.Fatal(err)
	}
	_, err = device.Scan()
	if err != nil {
		t.Fatal(err)
	}
}
