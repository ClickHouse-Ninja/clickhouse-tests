package tests

import (
	"testing"

	"github.com/ClickHouse-Ninja/ok"
	"github.com/stretchr/testify/assert"
)

func TestSelect10FromSystemNumbers(t *testing.T) {
	ok := ok.Connect(t, "tcp://127.0.0.1:9000?debug=0")
	if ok.TableExists("system", "numbers") {
		var result []int
		if rows, err := ok.DB().Query("SELECT number FROM system.numbers LIMIT 10"); assert.NoError(t, err) {
			defer rows.Close()
			for rows.Next() {
				var i int
				if err := rows.Scan(&i); !assert.NoError(t, err) {
					return
				}
				result = append(result, i)
			}
			if assert.Len(t, result, 10) {
				var expected []int
				for i := 0; i < 10; i++ {
					expected = append(expected, i)
				}
				assert.Equal(t, expected, result)
			}
		}
	}
}
