package goutils

// Deep copy input byte slice.
func DeepCopyBytes(b []byte) []byte {
	rb := make([]byte, len(b))
	copy(rb, b)
	return rb
}
