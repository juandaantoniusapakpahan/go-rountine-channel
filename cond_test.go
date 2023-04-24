package gorountinechannel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var group = sync.WaitGroup{}

func WaitCondition(number int) {
	defer group.Done()
	group.Add(1)
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done: ", number)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {

	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for j := 0; j < 10; j++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()
	group.Wait()

}
