package environment

import (
	"os"

	"github.com/leviatan89/api.victorsesma.com/types"
)

// GetCofniguration holds the app environment variables
func GetCofniguration() types.Configuration {
	return types.Configuration{
		SSLStatus:     os.Getenv("SSL_ON"),
		Privkey:       os.Getenv("PRIVATE_KEY"),
		Fullchain:     os.Getenv("FULL_CHAIN"),
		DBDSN:         os.Getenv("DB_DSN"),
		HTTPPort:      os.Getenv("HTTP_PORT"),
		HTTPSPort:     os.Getenv("HTTPS_PORT"),
		MigrationsDir: os.Getenv("MIGRATIONS_DIR"),
		DBDSNTest:     os.Getenv("DBDSN_TEST"),
	}
}
