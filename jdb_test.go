package jdb_test

import (
	"testing"
	"github.com/puresoul/jdb"
)

func TestOpen(t *testing.T) {
	jdb := jdb.Open("test.db")

	jdb.Map["test1"] = "test2"

	jdb.Close()
}

func TestClose(t *testing.T) {
	jdb := jdb.Open("test.db")

	if jdb.Map["test1"] != "test2" {
		t.Error("LoadConfig should have failed")
	}

	jdb.Close()
}
