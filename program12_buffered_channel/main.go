package main

/*
Channels in Go are a powerful feature that allows goroutines to communicate with each other in order to synchronize
 their execution and share data. Think of channels as pipes through which you can send and receive values with the
 channel operator (<-). By design, they ensure that data exchanges between goroutines are synchronized, making concurrent
  programming in Go more accessible and safer.
  You use the channel operator <- to send and receive data:
  ch <- v  // Send v to channel ch
  v := <-ch // Receive from ch, and assign value to v
 In buffered channels, sends are only blocked when the buffer is full, and receives are blocked when the buffer is empty.
 Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression:
v, ok := <-ch
Here, ok is false if no value is received because the channel is closed and empty.
You can also use a for loop with the range keyword to receive values from a channel until it is closed:
Copy code
for v := range ch {
    fmt.Println(v)
}
The select statement lets a goroutine wait on multiple communication operations. A select blocks until one of its cases can run,
 then it executes that case. It's similar to a switch statement but for channels:
select {
case v := <-ch1:
    fmt.Println("Received from ch1", v)
case ch2 <- 42:
    fmt.Println("Sent 42 to ch2")
default:
    fmt.Println("Neither send nor receive operations were ready")
}
The select statement lets a goroutine wait on multiple communication operations. A select blocks until one of its cases can run, then it executes that case. It's similar to a switch statement but for channels:

go
Copy code
select {
case v := <-ch1:
    fmt.Println("Received from ch1", v)
case ch2 <- 42:
    fmt.Println("Sent 42 to ch2")
default:
    fmt.Println("Neither send nor receive operations were ready")
}
*/

// package main

import (
	"fmt"
)
// The producer is not blocked until the length of the channel is 5. The consumer reads 5 values continuously, and then 
//it is blocked because the channel is empty.
// NB !!!!!!!!!!!!! STH IS WRONG HERE:THE OUTPUT IS NOT AS EXPECTED: PROGRAM 11.15: BUFFERED

func main() {
	ch := make(chan int, 5)
	done := make(chan bool)
	go producer(ch, done)
	go consumer(ch) // reads from producer
	<-done // wait intil go routine producer to finish, because it is the onlye one that sends a value on the done channel!! 
}

func producer(ch chan int, done chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("PRODUCER: sending %v\n", i)
		ch <- i // sending i through the channel
	}
	close(ch)
	done <- true // send true through the done channel, done is not closed!
}

func consumer(ch chan int) {
	for val := range ch {
		fmt.Printf("CONSUMER: read %v\n", val)
	}
}

/*
func readVal(ch <-chan int) {  // compilation error: the ch variable is declared as a read-only variable.
 ch <- 10
}

func writeVal(ch chan<- int) {
 <-ch // code will not compile because we are trying to read from the send-only channel.
}

*/
