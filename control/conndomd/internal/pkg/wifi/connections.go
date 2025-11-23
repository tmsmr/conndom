package wifi

import "github.com/google/uuid"

type ConfiguredNetwork struct {
	ID   string
	UUID string
	SSID string
}

func (m *Manager) ListConfiguredNetworks() ([]ConfiguredNetwork, error) {
	conns, err := m.settings.ListConnections()
	if err != nil {
		return nil, err
	}
	var networks []ConfiguredNetwork
	for _, conn := range conns {
		settings, err := conn.GetSettings()
		if err != nil {
			return nil, err
		}
		if settings["connection"]["type"] != "802-11-wireless" || settings["802-11-wireless"]["mode"] != "infrastructure" {
			continue
		}
		networks = append(networks, ConfiguredNetwork{
			ID:   settings["connection"]["id"].(string),
			UUID: settings["connection"]["uuid"].(string),
			SSID: string(settings["802-11-wireless"]["ssid"].([]uint8)),
		})
	}
	return networks, nil
}

func (m *Manager) AddConfiguredNetwork(ssid string, passphrase string) (*ConfiguredNetwork, error) {
	settings := map[string]map[string]interface{}{
		"connection": {
			"id":             ssid,
			"uuid":           uuid.NewString(),
			"type":           "802-11-wireless",
			"interface-name": m.conf.StaWifiInterface,
			"autoconnect":    true,
		},
		"802-11-wireless": {
			"mode": "infrastructure",
			"ssid": []uint8(ssid),
		},
		"ipv4": {
			"method": "auto",
		},
		"ipv6": {
			"method": "ignore",
		},
	}
	if passphrase != "" {
		settings["802-11-wireless-security"] = map[string]interface{}{
			"key-mgmt": "wpa-psk",
			"psk":      passphrase,
		}
	}
	conn, err := m.settings.AddConnection(settings)
	if err != nil {
		return nil, err
	}
	settingsConn, err := conn.GetSettings()
	if err != nil {
		return nil, err
	}
	return &ConfiguredNetwork{
		ID:   settingsConn["connection"]["id"].(string),
		UUID: settingsConn["connection"]["uuid"].(string),
		SSID: string(settingsConn["802-11-wireless"]["ssid"].([]uint8)),
	}, nil
}

func (m *Manager) DeleteConfiguredNetwork(uuid string) error {
	conn, err := m.settings.GetConnectionByUUID(uuid)
	if err != nil {
		return err
	}
	err = conn.Delete()
	if err != nil {
		return err
	}
	return nil
}
