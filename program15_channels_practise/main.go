package main

import (
	"fmt"
	"time"
)

/*
Write a producer-consumer program where there are two producer goroutines that produce data int at different intervals. One goroutine will produce data
 every 300 milliseconds and another goroutine will produce data every 400 milliseconds. There will be one consumer goroutine that will consume the data.
  Please note that producers should produce data 100 times only.
*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1, 300*time.Millisecond) // send ints to ch1 and ch2
	go producer(ch2, 400*time.Millisecond)

	go consumer(ch1, ch2) // recieve ints from ch1 nad ch2

	time.Sleep(45 * time.Second)
}

func producer(ch chan int, interval time.Duration) {
	for i := 0; i < 100; i++ {
		time.Sleep(interval)
		ch <- i
	}
	close(ch)
}

func consumer(ch1, ch2 chan int) {
	for {
		select {
		case val, ok := <-ch1:
			if ok {
				fmt.Printf("Received %d from producer1\n", val)
			} else {
				ch1 = nil
			}
		case val, ok := <-ch2:
			if ok {
				fmt.Printf("Received %d from producer2\n", val)
			} else {
				ch2 = nil
			}
		}
		if ch1 == nil && ch2 == nil {
			fmt.Println("All producers are done producing.")
			return
		}
	}
}

// find the error: package main

/* import (
	"fmt"
	"sync"
  )

  func main() {
   wg := sync.WaitGroup{}
   wg.Add(2) // should be 1, we use only 1 go routine
   go func() {
	fmt.Println("Running go routine")
	wg.Done()
   }()
   wg.Wait()
  } */

// what is the difference between buffered and unbuffered channels:

/*
  Asynchronous: Buffered channels are asynchronous; sending or receiving data on a buffered channel will not block
   if the channel buffer is not full (on send) or not empty (on receive).
Has Capacity: They have a capacity you specify when you create the channel. The channel can hold a limited number
 of values without a corresponding receiver for those values.
No Immediate Handoff: They don't guarantee that the other side of the channel is ready to receive/send data at
 the moment you perform a send/receive operation. The operation will only block when the buffer is full (on send)
  or empty (on receive).
Example of a buffered channel:

go
Copy code
ch := make(chan int, 10) // buffered channel with capacity of 10
In summary, the choice between buffered and unbuffered channels affects the behavior and synchronization of your
 Go program. Unbuffered channels are useful when you want to synchronize goroutines (for example, to ensure that
	a goroutine has finished processing before the next step), while buffered channels are useful when you want
	to allow goroutines to proceed without immediate synchronization, up to the limit of the buffer.
*/

/*
 Which of the following code defines a read-only channel as a function argument?

 func read(ch <-chan int) { // receiver only channel
  ch <- 10
 }
 func read(ch chan<- int) { // send only channel
  ch <- 10
 }
 // this one is read only
 func read(ch <-chan int) {
  value := <-ch // This is a correct usage of a receive-only channel
  // ... use value ...
}

!!! NB !!!!
chan int is a bidirectional channel.
<-chan int is a receive-only channel (read-only).
chan<- int is a send-only channel (write-only).
*/

/*
 What will happen when data is sent to an unbuffered channel but there is no goroutine to read that data?

The program will be stuck for an infinite time.

The program will panic because it is a deadlock.

The program will end whether or not there is a listener.

--- ANSWER: -----------------

When data is sent to an unbuffered channel and there is no goroutine to read that data, the sending goroutine
 will block until another goroutine reads from the channel. If there is never a goroutine that comes to read the
  data, the sending goroutine will remain blocked indefinitely. This can lead to a deadlock if no other activity
   in the program can proceed, and if the Go runtime detects that all goroutines are deadlocked, it will panic 
   with a deadlock error message.

So, if it's the only activity left in the program and there are no other goroutines that will ever be scheduled
 to read from the channel, then:
The program will panic because it is a deadlock.
*/
