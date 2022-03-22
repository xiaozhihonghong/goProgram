package geeCache

type ByteView struct {
	B []byte
}

func (vb ByteView) Len() int {
	return len(vb.B)
}

func Copy(b []byte) []byte  {
	t := make([]byte, len(b))
	copy(t, b)
	return t
}
