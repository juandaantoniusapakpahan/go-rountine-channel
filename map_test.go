package gorountinechannel

import (
	"fmt"
	"sync"
	"testing"
)

func MapKeyValueAdd(mapTest *sync.Map, keyVal int, group *sync.WaitGroup) {
	mutex := &sync.Mutex{}

	defer group.Done()
	mutex.Lock()
	group.Add(1)
	mutex.Unlock()
	mapTest.Store(keyVal, keyVal)
}

func TestMap(t *testing.T) {
	group := &sync.WaitGroup{}
	mapTest := &sync.Map{}

	for i := 0; i < 100; i++ {
		go MapKeyValueAdd(mapTest, i, group)
	}
	group.Wait()

	mapTest.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
