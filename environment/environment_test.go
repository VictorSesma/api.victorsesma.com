package environment

import (
	"os"
	"testing"

	"github.com/leviatan89/api.victorsesma.com/types"
	"github.com/stretchr/testify/assert"
)

func TestGetCofniguration(t *testing.T) {
	os.Setenv("SSL_ON", "true")
	os.Setenv("PRIVATE_KEY", "/I/am/here")
	os.Setenv("FULL_CHAIN", "/I/am/here/2")
	os.Setenv("DB_DSN", "cool_db_dsn")
	os.Setenv("HTTP_PORT", "cool_http_port")
	os.Setenv("HTTPS_PORT", "cool_https_port")
	os.Setenv("MIGRATIONS_DIR", "./cool_migrations_dir")
	os.Setenv("DBDSN_TEST", "cool_db_dsn_test")

	expected := types.Configuration{
		SSLStatus:     "true",
		Privkey:       "/I/am/here",
		Fullchain:     "/I/am/here/2",
		DBDSN:         "cool_db_dsn",
		HTTPPort:      "cool_http_port",
		HTTPSPort:     "cool_https_port",
		MigrationsDir: "./cool_migrations_dir",
		DBDSNTest:     "cool_db_dsn_test",
	}

	conf := GetCofniguration()
	assert.Equal(t, expected, conf)
}
