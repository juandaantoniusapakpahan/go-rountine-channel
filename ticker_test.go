package gorountinechannel

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(5 * time.Second)

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

func TestTickerStop(t *testing.T) {
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		time.Sleep(10 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

func TestTickerStopSelect(t *testing.T) {
	ticker := time.NewTicker(5 * time.Second)

	select {
	case data := <-ticker.C:
		fmt.Println(data)
	case <-time.After(11 * time.Second):
		ticker.Stop()
	}
}

func TestTick(t *testing.T) {
	ticker := time.Tick(2 * time.Second)

	for data := range ticker {
		fmt.Println(data)
	}
}
