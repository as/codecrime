// This program "misuses" unsafe.Pointer twice, but go vet
// only prints one warning.
package main

import (
	"unsafe"
)

var four byte = 4

func main() {
	p := uintptr(unsafe.Pointer(&four))
	v := *((*byte)(unsafe.Pointer(p)))
	w := **((**byte)(unsafe.Pointer(&p))) // hahaha =(
	println(p,v,w,&four)
}

