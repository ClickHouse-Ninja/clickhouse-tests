package tests

import (
	"bytes"
	"encoding/csv"
	"testing"

	"github.com/ClickHouse-Ninja/ok"
	"github.com/stretchr/testify/assert"
)

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

func TestArgMax(t *testing.T) {
	const ddl = `
	CREATE DATABASE test;
	CREATE TABLE    test.arg_max(
		user_id Int8
		, value Int8
	) Engine Memory
	`
	var (
		ok = ok.Connect(t, "tcp://127.0.0.1:9000?debug=0")
	)
	if ok.DatabaseExists("test") {
		t.Fatal("database 'test' is already exists")
	}
	defer ok.Clear()
	if err := ok.Exec(ddl); assert.NoError(t, err) {
		var (
			buf bytes.Buffer
			csv = csv.NewWriter(&buf)
		)
		csv.WriteAll([][]string{
			[]string{"24", "25"},
			[]string{"42", "52"},
		})
		csv.Flush()

		if ok.CopyFromCSVReader(&buf, "INSERT INTO test.arg_max VALUES") {
			var value int
			if err := ok.DB().QueryRow("SELECT argMax(user_id, value) FROM test.arg_max").Scan(&value); assert.NoError(t, err) {
				assert.Equal(t, 42, value)
			}
		}
	}
}
