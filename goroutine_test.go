package gorountinechannel

import (
	"fmt"
	"testing"
	"time"
)

func SayHello(number int) {
	fmt.Println("Hello ", number)
}

func TestGoRoutine(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go SayHello(i)
	}
	fmt.Println("Have finished ?")
	time.Sleep(5 * time.Second)
}

func TestChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello guys, My name is Chiky"
	}()

	data := <-channel
	fmt.Println(data)
	fmt.Println("Done")
	time.Sleep(3 * time.Second)

}

func ChannelAsParam(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello guys, this is My Wifi"
}

func TestChannelAsParam(t *testing.T) {
	channel := make(chan string)

	go ChannelAsParam(channel)

	data := <-channel
	fmt.Println(data)
	fmt.Println("Done")

}
