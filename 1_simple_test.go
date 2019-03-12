package tests

import (
	"testing"

	"github.com/ClickHouse-Ninja/ok"
	"github.com/stretchr/testify/assert"
)

func TestSelect1(t *testing.T) {
	var (
		value int
		ok    = ok.Connect(t, "tcp://127.0.0.1:9000?debug=0")
	)
	if err := ok.DB().QueryRow("SELECT 1").Scan(&value); assert.NoError(t, err) {
		assert.Equal(t, 1, value)
	}
}
