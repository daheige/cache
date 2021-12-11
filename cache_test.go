package cache

import (
	"log"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cacheEntry, err := New(10 * time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	cacheEntry.Set("abc", []byte("123"))

	time.Sleep(2 * time.Second)
	log.Println(cacheEntry.Get("abc"))

	cacheEntry.SetJson("abc", []string{"a", "b"})

	time.Sleep(2 * time.Second)

	bean := []string{}
	log.Println(cacheEntry.GetJson("abc", &bean))

	log.Println("bean: ", bean)

	cacheEntry.SetJson("my_map", map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": "abc",
	})

	time.Sleep(1 * time.Second)
	m := map[string]interface{}{}
	cacheEntry.GetJson("my_map", &m)

	log.Println("m: ", m)
}

/*
=== RUN   TestNew
2021/12/11 22:06:01 [49 50 51] <nil>
2021/12/11 22:06:03 <nil>
2021/12/11 22:06:03 bean:  [a b]
2021/12/11 22:06:04 m:  map[a:1 b:2 c:abc]
--- PASS: TestNew (5.03s)
PASS
*/
