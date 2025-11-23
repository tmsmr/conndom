package wifi

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
