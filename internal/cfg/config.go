package cfg

import (
	"fmt"
	"os"
	"sync"
)

type Cfg struct {
	Port     string
	PgUrl    string
	TokenKey string
}

var (
	once   sync.Once
	config *Cfg
)

func GetConfig() *Cfg {
	once.Do(func() {
		config = &Cfg{}
		config.Port = os.Getenv("APP_PORT")

		config.PgUrl = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			os.Getenv("PG_USER"),
			os.Getenv("PG_PASSWORD"),
			os.Getenv("PG_HOST"),
			os.Getenv("PG_PORT"),
			os.Getenv("PG_DB_NAME"),
			os.Getenv("PG_SSL_MODE"),
		)

		config.TokenKey = os.Getenv("TOKEN_KEY")
	})

	return config
}
