package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	_ "unsafe"
)

//go:linkname x image.formats

var x []uintptr

func main() {
	fmt.Println(image.Decode(bytes.NewReader([]byte{137, 80, 78, 71, 13, 10, 26, 10})))
	x = nil
	fmt.Println(image.Decode(bytes.NewReader([]byte{137, 80, 78, 71, 13, 10, 26, 10})))
}

// output
// <nil> png unexpected EOF
// <nil>  image: unknown format
