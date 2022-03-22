package geeCache

import (
	"fmt"
	"log"
	"testing"
)

var db = map[string]string{
	"chenzhi": "89",
	"yangjin": "79",
	"wangtao": "85",
}

func TestGroup_GetCache(t *testing.T) {
	loadCount := make(map[string]int, len(db))

	gee := NewGroup("scores", 1000000, GetterFunc(
			func(key string)([]byte, error){
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				if _, ok := loadCount[key]; !ok {
					loadCount[key] = 0
				} else {     // 这里需要加上else
					loadCount[key] += 1
				}
				return []byte(v), nil
			}
			return nil, fmt.Errorf("DB is not exist")
			}))
	for k, v := range db {
		if bv, err := gee.GetCache(k); err != nil || string(bv.b) != v {
			t.Fatal("failed to get value of Tom")
		}
		if _, err := gee.GetCache(k); err != nil || loadCount[k] > 1 {
			t.Fatal("miss cache")
		}
	}
	if view, err := gee.GetCache("unknown"); err == nil {
		t.Fatalf("the value of unknow should be empty, but %s got", view)
	}

}
