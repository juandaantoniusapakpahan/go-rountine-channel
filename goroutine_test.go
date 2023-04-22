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

func OnlyIn(channel chan<- string) {
	channel <- "Hello Channel In Out"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go ChannelAsParam(channel1)
	go ChannelAsParam(channel2)
	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Channel 2: ", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}
