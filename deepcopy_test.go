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
