package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	NewDecoder() // JSON的流式读写
}

func NewDecoder() {
	// func NewDecoder(r io.Reader) *Decoder {
	dec := json.NewDecoder(os.Stdin)
	// func NewEncoder(w io.Writer) *Encoder {
	enc := json.NewEncoder(os.Stdout)
	fmt.Println("dec = ", dec)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Title" {
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
