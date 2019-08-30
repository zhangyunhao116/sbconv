package goutils

func DeepCopyString(s string) string {
	return string(StringToBytes(s))
}

func DeepCopyBytes(b []byte) []byte {
	rb := make([]byte, len(b))
	copy(rb, b)
	return rb
}
