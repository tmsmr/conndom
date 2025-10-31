package network

import (
	"os"
	"sync"

	gonm "github.com/Wifx/gonetworkmanager/v3"
	"github.com/tmsmr/conndom/conndomd/internal/pkg/env"
)

var once sync.Once

type Manager struct {
	nm gonm.NetworkManager
}

func NewManager(conf env.Config) (*Manager, error) {
	once.Do(func() {
		if err := os.Setenv(DbusSystemBusAddrEnv, conf.DbusSystemBusAddress); err != nil {
			panic(err)
		}
	})
	nm, err := gonm.NewNetworkManager()
	if err != nil {
		return nil, err
	}
	return &Manager{
		nm: nm,
	}, nil
}
