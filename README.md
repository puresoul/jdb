# jdb
Compressed json flatfile "database" (it's golang map datatype actualy) with simple key/value api

Example how to open (one time file read), read from db and save it to file

```go
package main

import (
	"github.com/puresoul/jdb"
	"fmt"
)

func main() {
	// Create database file
	d := jdb.Open("test.db")

	// Insert string
	d.Map["test"] = "test"
	fmt.Println(d.Map)

	// Save
	_ = d.Close()

	d = jdb.Open("test.db")

	// Read String
	t := j.ReadStr("test")
	fmt.Println(tst)

	// Write Int
	d.Map["tst"] = 123
	_ = d.Close()

	d = jdb.Open("test.db")

	// Read Int
	i := j.ReadInt("tst")
	fmt.Println(i+i)
}
``
