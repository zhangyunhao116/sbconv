package sbconv

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

func TestStringToBytesNew(t *testing.T) {
	s := "TEST"
	sh := *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	b := StringToBytesNew(s)
	bh := *(*reflect.StringHeader)(unsafe.Pointer(&b))
	if sh.Data == bh.Data {
		t.Error("got previous data")
	}
	b[0] = 65
	if BytesToString(b) != "AEST" {
		t.Error("can not change returned byte slice")
	}
}

func TestBytesToStringNew(t *testing.T) {
	b := []byte{65, 66, 67, 68, 69}
	bh := *(*reflect.StringHeader)(unsafe.Pointer(&b))
	s := BytesToStringNew(b)
	sh := *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	if sh.Data == bh.Data {
		t.Error("got previous data")
	}
}
