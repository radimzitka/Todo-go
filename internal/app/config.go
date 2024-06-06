package app

import "github.com/spf13/viper"

type Config struct {
	Net struct {
		Port string
	}
	Db struct {
		ConnectionString string
	}
	JWTSecret string
}

func LoadConfig(filepath string, cfg *Config) error {
	viper.SetConfigFile(filepath)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(cfg)
	return err
}
