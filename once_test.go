package gorountinechannel

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0
var noOne = 0

func OnceCount() {
	counter++
}
func MoreThanOne() {
	noOne++
}

func TestOnce(t *testing.T) {
	once := &sync.Once{}
	group := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnceCount)
			mutex.Lock()
			MoreThanOne()
			mutex.Unlock()
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Once Counter :", counter)
	fmt.Println("Mutex: ", noOne)
}
