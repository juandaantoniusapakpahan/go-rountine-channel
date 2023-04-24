package gorountinechannel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Start", time.Now())

	timeDur := <-timer.C

	fmt.Println("End", timeDur)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("Start", time.Now())

	timeDura := <-channel
	fmt.Println("End", timeDura)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("End", time.Now())
		group.Done()
	})
	fmt.Println("Start", time.Now())
	group.Wait()
}
