package geeCache

import (
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	f := GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})
	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(expect, v) {
		t.Errorf("callBack error")
	}
}
