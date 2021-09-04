package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Entry struct {
	Key string
	Val string
}

func main() {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	e := Entry{"k1", "v1"}
	enc.Encode(e)

	var ent Entry
	dec.Decode(&ent)

	fmt.Println(buf.Bytes())
	fmt.Println(ent.Key, ent.Val)
}
