package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1 = make(chan string)
	ch2 := make(chan string)
	go producer(ch1, "P1", time.Millisecond*1000) // turn into seconds
	go producer(ch2, "P2", time.Millisecond*1000)
	receiver(ch1, ch2)
}

func producer(ch chan string, name string, sleep time.Duration) {
	for i := 0; i < 10; i++ {
		time.Sleep(sleep)
		ch <- fmt.Sprintf("%v: %v", name, i)
	}
	close(ch)
}

func receiver(ch1 chan string, ch2 chan string) {
	for {
		select { //A select blocks until one of its cases can run, then it executes that case. If multiple cases can proceed, it picks one at random.
		case msg, ok := <-ch1: // receives messages
			fmt.Println("Message from ch1", msg)
			if !ok { return }
		case msg, ok := <-ch2:
			fmt.Println("Message from ch2", msg)
			if !ok { return }
		}
	}
}
/*
NB!!! --- When a channel is closed, you can still receive values from it. If there are no more values to be received, the second value returned by the 
receive operation (ok in this code) will be false. This is often used to detect when a channel has been closed and no more data will come through it.
*/