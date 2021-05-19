package jdb

import (
	"encoding/json"
	"io/ioutil"
	"compress/gzip"
	"os"
	"fmt"
	"errors"
)

type Jdb struct {
	Name string
	Map Db
}

type Json struct {
	Output []Data  `json:"output"`
}

type Data struct {
	Key string `json:"key"`
	Value interface{} `json:"value"`
}

type Db map[string]interface{}

func Open(f string) (*Jdb) {
	zipReader, err := os.Open(f)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, _ := os.Create(f)
			w := gzip.NewWriter(file)
			w.Write([]byte(`{"output":[]}`))
			zipReader, _ = os.Open(f)
		} else {
			fmt.Println("Database file is in use!")
			return &Jdb{}
		}
	}

	defer zipReader.Close()

	reader, err := gzip.NewReader(zipReader)
	if err != nil {
		panic(err)
	}

	var j Json

	d := &Jdb{Name: f, Map: make(map[string]interface{})}

	content, _ := ioutil.ReadAll(reader)

	_ = json.Unmarshal(content, &j)

	for _, v := range j.Output {
		d.Map[v.Key] = v.Value
	}

	return d
}

func (jdb *Jdb) ReadStr(key string) string {
	switch jdb.Map[key].(type) {
	case int:
		return fmt.Sprint(jdb.Map[key].(int))
	case float64:
		return fmt.Sprint(jdb.Map[key].(float64))
	}
	return jdb.Map[key].(string)
}

func (jdb *Jdb) ReadFloat(key string) float64 {
	switch jdb.Map[key].(type) {
	case int:
		return float64(jdb.Map[key].(int))
	}
	return jdb.Map[key].(float64)
}

func (jdb *Jdb) ReadInt(key string) int {
	switch jdb.Map[key].(type) {
	case float64:
		return int(jdb.Map[key].(float64)) 
	}
	return jdb.Map[key].(int)
}

func (jdb *Jdb) Close() error {
	var j Json

	for k, v := range jdb.Map {
		j.Output = append(j.Output, Data{Key: k, Value: v})
	}

	out, err := json.Marshal(j)

	if err != nil {
		return err
	}

	file, err := os.Create(jdb.Name+".bak")

	defer file.Close()

	if err != nil {
		return err
	}

	w := gzip.NewWriter(file)
	w.Write(out)
	w.Close()

	err = os.Rename(jdb.Name+".bak", jdb.Name)

	if err != nil {
		return err
	}

	return nil
}
