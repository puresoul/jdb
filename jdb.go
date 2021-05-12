package jdb

import (
	"encoding/json"
	"io/ioutil"
	"compress/gzip"
	"os"
)

type Jdb struct {
	Name string
	Map map[string]string
}

type Json struct {
	Output []Data  `json:"output"`
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
		zipReader, _ = os.Open(f)
	}

	defer zipReader.Close()

	reader, err := gzip.NewReader(zipReader)
	if err != nil {
		panic(err)
	}

	var j Json

	d := &Jdb{Name: f, Map: make(map[string]string)}

	content, _ := ioutil.ReadAll(reader)

	_ = json.Unmarshal(content, &j)

	for _, v := range j.Output {
		d.Map[v.Key] = v.Value
	}

	return d
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

	err = os.Remove(jdb.Name)

	if err != nil {
		return err
	}

	file, err := os.Create(jdb.Name)

	defer file.Close()

	if err != nil {
		return err
	}

	w := gzip.NewWriter(file)
	w.Write(out)
	w.Close()

	return nil
}
