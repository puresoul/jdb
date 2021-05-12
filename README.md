# jdb
Compressed json flatfile database with simple key/value api

Example how to open, read from db and save it to file

```go
package main

import (
	"github.com/puresoul/jdb"
	"fmt"
)

func main() {
	d := jdb.Open("test.db")
	d.Map["test1"] = "test2"
	fmt.Println(d.Map)
	_ = d.Close()
}
``
