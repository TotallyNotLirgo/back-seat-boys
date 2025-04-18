package general

import (
	"os"
	"reflect"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_ENV         string `default:"DEV"`
	TRUSTED_PROXIES string `default:"127.0.0.1"`
	PORT            string `default:":8080"`
	CERT_FILE       string `default:"cert/localhost.crt"`
	KEY_FILE        string `default:"cert/localhost.key"`
}

func initConfig() Config {
	_ = godotenv.Load()
	config := Config{}
	vPointer := reflect.ValueOf(&config)
	vType := vPointer.Type().Elem()
	vNew := reflect.New(vType).Elem()
	vNewType := vNew.Type()
	for i := range vNewType.NumField() {
		env, ok := os.LookupEnv(vNewType.Field(i).Name)
		if !ok {
			env = vNewType.Field(i).Tag.Get("default")
		}
		vNew.Field(i).Set(reflect.ValueOf(env))
	}
	vPointer.Elem().Set(vNew)
	return config
}

var (
	GetConfig func() Config = sync.OnceValue(initConfig)
)
