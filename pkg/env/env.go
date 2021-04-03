package env

import (
	"github.com/spf13/viper"
	"log"
)

func LoadEnvVar(key string) string {
	viper.AddConfigPath("../../")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error %s: while reading .env file using viper!\n", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalln("Invalid type assertion")

	}
	return value

}
