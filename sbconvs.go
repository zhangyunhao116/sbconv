package goutils

import (
	"reflect"
	"unsafe"
)

// Converts string to a byte slice.
//
// This is a shallow copy, means that the returned byte slice reuse
// the underlying array in string, so you can't change the returned
// byte slice in any situations.
func StringToBytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

// Converts string to a byte slice.
//
// This is a shallow copy, means that the returned string reuse the
// underlying array in byte slice, it's your responsibility to keep
// the input byte slice survive until you don't access the string anymore.
func BytesToString(b []byte) (s string) {
	return *(*string)(unsafe.Pointer(&b))
}

// Converts string to a byte slice.(Deep copy)
func StringToBytesNew(s string) (b []byte) {
	return DeepCopyBytes(StringToBytes(s))
}

// Converts byte slice to a string.(Deep copy)
func BytesToStringNew(b []byte) (s string) {
	return DeepCopyString(BytesToString(b))
}
