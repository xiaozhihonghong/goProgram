package cache

type ByteView struct {
	b []byte
}

func (vb ByteView) Len() int {
	return len(vb.b)
}

func Copy(b []byte) []byte  {
	t := make([]byte, len(b))
	copy(t, b)
	return t
}
