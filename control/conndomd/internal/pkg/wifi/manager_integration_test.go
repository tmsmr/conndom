package wifi_test

import (
	"testing"

	"github.com/tmsmr/conndom/conndomd/internal/pkg/env"
	"github.com/tmsmr/conndom/conndomd/internal/pkg/wifi"
)

func TestManager_GetStaDevice(t *testing.T) {
	conf, err := env.Load()
	if err != nil {
		t.Fatal(err)
	}
	m, err := wifi.NewManager(*conf)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.GetStaDevice()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDevice_Scan(t *testing.T) {
	conf, err := env.Load()
	if err != nil {
		t.Fatal(err)
	}
	m, err := wifi.NewManager(*conf)
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

func TestManager_ListConnections(t *testing.T) {
	conf, err := env.Load()
	if err != nil {
		t.Fatal(err)
	}
	m, err := wifi.NewManager(*conf)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.ListConfiguredNetworks()
}

func TestManager_AddDeleteConfiguredNetwork(t *testing.T) {
	conf, err := env.Load()
	if err != nil {
		t.Fatal(err)
	}
	m, err := wifi.NewManager(*conf)
	if err != nil {
		t.Fatal(err)
	}
	network, err := m.AddConfiguredNetwork("test-ssid", "test-passphrase")
	if err != nil {
		t.Fatal(err)
	}
	err = m.DeleteConfiguredNetwork(network.UUID)
	if err != nil {
		t.Fatal(err)
	}
}
