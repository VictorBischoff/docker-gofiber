package utilities

import "github.com/spf13/viper"

type Config struct {
	POSTGRES_USER string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_HOST string `mapstructure:"POSTGRES_HOST"`
	DSN string `mapstructure:"DSN"`
}

func LoadConfigPath(path string) (config Config, err error){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)
    return
}