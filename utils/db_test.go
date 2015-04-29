package utils

import (
	"testing"
)

func TestGetConnection(t *testing.T) {
	db := GetConnection()

	err := db.Ping()
	if err != nil {
		t.Errorf("GetConnection() failed.")
	}
}
