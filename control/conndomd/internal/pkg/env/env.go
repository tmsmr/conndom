package env

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbusSystemBusAddress string `split_words:"true"`
}

func Load(prefix string) (*Config, error) {
	var e Config
	err := envconfig.Process(prefix, &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
