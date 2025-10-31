package network_test

import (
	"testing"

	"github.com/tmsmr/conndom/conndomd/internal/pkg/env"
	"github.com/tmsmr/conndom/conndomd/internal/pkg/network"
)

const forwardedDbusSocket = "/tmp/conndom_dbus_system.sock"

func TestManager_ListWifiDevices(t *testing.T) {
	m, err := network.NewManager(env.Config{
		DbusSystemBusAddress: forwardedDbusSocket,
	})
	if err != nil {
		t.Fatalf("NewManager() error = %v", err)
	}
	devices, err := m.ListWifiDevices()
	if err != nil {
		t.Fatalf("ListWifiDevices() error = %v", err)
	}
	for _, d := range devices {
		t.Logf("%v\n", d)
	}
}
