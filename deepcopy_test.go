package goutils

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestDeepCopyBytes(t *testing.T) {
	var b = []byte("11111")
	var bc = DeepCopyBytes(b)

	if string(b) != string(bc) {
		t.Fatal("error content")
	}

	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bch := (*reflect.SliceHeader)(unsafe.Pointer(&bc))
	if bh.Data == bch.Data || bh.Len != bch.Len {
		t.Fatal("error headers")
	}

}

func TestDeepCopyString(t *testing.T) {
	var s = "111"
	var sc = DeepCopyString(s)

	if s != sc {
		t.Fatal("error content")
	}

	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sch := (*reflect.StringHeader)(unsafe.Pointer(&sc))
	if sh.Data == sch.Data || sh.Len != sch.Len {
		t.Fatal("error headers")
	}

}
