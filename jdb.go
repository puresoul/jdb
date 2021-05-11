package jdb

import (
	"encoding/json"
	"io/ioutil"
	"compress/gzip"
	"os"
)

type Jdb struct {
	Name string
	Output []byte
}

type Json struct {
	Output []Data `json:"output"`
}

type Data struct {
	Key string `json:"key"`
	Value string `json:"value"`
}


type Map map[string]string

func Open(f string) (*Jdb) {
	zipReader, err := os.Open(f)

	if(err != nil){
		file, _ := os.Create(f)
		w := gzip.NewWriter(file)
		w.Write([]byte(`{"output":[]}`))
		w.Close()
		zipReader, _ = os.Open(f)
	}

	defer zipReader.Close()

	reader, err := gzip.NewReader(zipReader)
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(reader)

	return &Jdb{Name: f, Output: content}
}

func (jdb *Jdb) Read() (Map) {
	d := make(map[string]string)
	var j Json

	_ = json.Unmarshal(jdb.Output, &j)

	for _, v := range j.Output {
		d[v.Key] = v.Value
	}

	return d
}

func (jdb *Jdb) Write(key, value string) error {
	var j Json
	err := json.Unmarshal(jdb.Output, &j)

	if err != nil {
		return err
	}

	j.Output = append(j.Output, Data{Key: key, Value: value})

	out, err := json.Marshal(j)

	if err != nil {
		return err
	}

	f, err := os.Create(jdb.Name)

	if err != nil {
		return err
	}

	w := gzip.NewWriter(f)
	w.Write(out)
	w.Close()

	return nil
}

