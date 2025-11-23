package env

import (
	"github.com/kelseyhightower/envconfig"
)

type Spec struct {
	DbusSystemBusAddress string `split_words:"true"`

	ApWifiInterface string `split_words:"true" default:"wlan0"`
	ApSSID          string `split_words:"true" default:"CONNDOM"`
	ApPassphrase    string `split_words:"true" default:"insecure"`

	StaWifiInterface string `split_words:"true" default:"wlan1"`
}

func Load() (*Spec, error) {
	var e Spec
	err := envconfig.Process("", &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
