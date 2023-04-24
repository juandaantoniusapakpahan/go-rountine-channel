package gorountinechannel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyncronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("Koniciwa")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	waitgroup := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncronous(waitgroup)
	}
	waitgroup.Wait()
	fmt.Println("Completed")
}
