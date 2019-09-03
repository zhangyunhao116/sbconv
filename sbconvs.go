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

// Converts string to a byte slice. (Deep copy)
//
// The returned byte slice can be modified, it is actually a new byte slice,
// and it's content equal to the input string.
func StringToBytesNew(s string) []byte {
	b := make([]byte, len(s))

	// Inline, equal to `temp := StringToBytes(s)`
	var temp []byte
	tempH := (*reflect.SliceHeader)(unsafe.Pointer(&temp))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	tempH.Data = sh.Data
	tempH.Len = sh.Len
	tempH.Cap = sh.Len

	copy(b, temp)
	return b
}

// Converts byte slice to a string. (Deep copy)
//
// This function is used in case you want get a fresh version of string
// converted from the input byte slice, the returned string is a real
// string, this means the data field on this string is read only.
func BytesToStringNew(b []byte) (s string) {
	return string(b)
}
