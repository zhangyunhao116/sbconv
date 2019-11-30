package sbconv

import (
	"unsafe"
)

// DeepCopyBytes returns a deep copy of the input byte slice.
func DeepCopyBytes(b []byte) []byte {
	rb := make([]byte, len(b))
	copy(rb, b)
	return rb
}

// DeepCopyString returns a deep copy of the input string.
//
// This function is used in case of the input string
// is converted from BytesToString, this means it is
// not a real string, but share an underlying array with
// a byte slice, use this function to create a real string.
func DeepCopyString(s string) string {
	return string(*(*[]byte)(unsafe.Pointer(&s)))
}
