package goutils

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestBytesToString(t *testing.T) {
	var b = []byte("11111")
	var s = BytesToString(b)
	if s != "11111" {
		t.Fatal("error content")
	}

	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
	if sh.Data != bh.Data || sh.Len != bh.Len {
		t.Fatal("error headers")
	}

}

func TestStringToBytes(t *testing.T) {
	var s = "11111"
	var b = StringToBytes(s)
	if string(b) != "11111" {
		t.Fatal("error content")
	}

	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := *(*reflect.SliceHeader)(unsafe.Pointer(&b))
	if sh.Data != bh.Data || sh.Len != bh.Len || sh.Len != bh.Cap {
		t.Fatal("error headers")
	}

}
