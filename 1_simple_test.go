package tests

import (
	"testing"

	"github.com/ClickHouse-Ninja/ok"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	if version, err := ok.Connect(t, "tcp://127.0.0.1:9000?debug=0").Version(); assert.NoError(t, err) {
		t.Logf("ClickHouse server version: %s", version)
	}
}
func TestSelect1(t *testing.T) {
	var (
		value int
		ok    = ok.Connect(t, "tcp://127.0.0.1:9000?debug=0")
	)
	if ok.DatabaseExists("system") {
		if err := ok.DB().QueryRow("SELECT 1").Scan(&value); assert.NoError(t, err) {
			assert.Equal(t, 1, value)
		}
	}
}

func TestDatabaseAndTableExists(t *testing.T) {
	ok := ok.Connect(t, "tcp://127.0.0.1:9000?debug=0")
	if assert.True(t, ok.DatabaseExists("system")) {
		for _, table := range []string{"clusters", "databases", "functions", "parts", "tables"} {
			assert.True(t, ok.TableExists("system", table))
		}
	}
}
