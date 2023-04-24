package gorountinechannel

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSpecCheck(t *testing.T) {
	cpu := runtime.NumCPU()
	fmt.Println("CPU:", cpu)
	thread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread:", thread)
	goroutine := runtime.NumGoroutine()
	fmt.Println("GoRun:", goroutine)
}
