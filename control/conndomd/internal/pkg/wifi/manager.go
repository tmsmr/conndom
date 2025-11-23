package wifi

import (
	"os"
	"sync"

	gonm "github.com/Wifx/gonetworkmanager/v3"
	"github.com/tmsmr/conndom/conndomd/internal/pkg/env"
)

var once sync.Once

type Manager struct {
	nm       gonm.NetworkManager
	settings gonm.Settings
	conf     env.Spec
}

func NewManager(conf env.Spec) (*Manager, error) {
	once.Do(func() {
		if err := os.Setenv(DbusSystemBusAddrEnv, conf.DbusSystemBusAddress); err != nil {
			panic(err)
		}
	})
	nm, err := gonm.NewNetworkManager()
	if err != nil {
		return nil, err
	}
	settings, err := gonm.NewSettings()
	if err != nil {
		return nil, err
	}
	return &Manager{
		nm:       nm,
		settings: settings,
		conf:     conf,
	}, nil
}

func (m *Manager) Running() bool {
	state, err := m.nm.State()
	if err != nil {
		return false
	}
	return state != gonm.NmStateUnknown
}

func (m *Manager) Reload() error {
	return m.nm.Reload(0)
}
