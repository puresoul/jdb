package jdb_test

import (
	"testing"
	"github.com/puresoul/jdb"
)

var (
	j *jdb.Jdb
)

func TestOpen(t *testing.T) {
	j = jdb.Open("test.db")
}

func TestWrite(t *testing.T) {
	j.Map["test"] = "test"

	j.Close()
}

func TestReadStr(t *testing.T) {
	j = jdb.Open("test.db")

	tst := j.ReadStr("test")

	if j.Map["test"] != tst {
		t.Error("This should never happen!")
	}
}

func TestClose(t *testing.T) {
	j.Map["int"] = 123
	j.Close()
}

func TestReadInt(t *testing.T) {
	j = jdb.Open("test.db")

	tst := j.ReadInt("int")

	if tst+tst != 123+123 {
		t.Error("This should never happen!")
	}

	s := j.ReadStr("int")

	if s != "123" {
		t.Error("This should never happen!")
	}

}
