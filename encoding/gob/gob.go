package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func main() {
	GobEncodeDecode()
}

type Message2 struct {
	Id   uint64
	Size uint64
	Data string
}

func GobEncodeDecode() {
	m1 := Message2{2, 1024, "gob"}
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	if err := enc.Encode(m1); err != nil {
		log.Fatal("encode error:", err)
	}
	fmt.Println(buf.Bytes())
	fmt.Println(m1)

	var m2 Message2
	if err := dec.Decode(&m2); err != nil {
		log.Fatal("decode error:", err)
	}
}
