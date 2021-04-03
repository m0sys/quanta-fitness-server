package config

import (
	"github.com/spf13/viper"
	"log"
)

type Confs struct {
	Database DatabaseConfs
}

type DatabaseConfs struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBTest     string
	DBDev      string
}

func LoadConfg(path string) Confs {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var confs Confs

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error %s: while reading .env file using viper!\n", err)
	}

	err = viper.Unmarshal(&confs)
	if err != nil {
		log.Fatalf("Error %s: while unmarshalling configurations!\n", err)
	}

	return confs
}
