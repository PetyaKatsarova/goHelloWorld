package main

import "fmt"

// import (
// 	"fmt"
// 	"sync"
// )

/*
The implementation of concurrency using channels: Do not communicate by sharing memory; instead, share memory by communicating.
*/

// func main() {
// 	ch := make(chan int)
// 	go writeVal(ch) // read value in the main function will be blocked until the wrtieVal function writes in the channel.
// 	// the main function will be waiting for the value to be available on the channel.
// 	/*
// 	Since read and write operations are blocking calls in the above code, if the main function is waiting for the read from the
// 	 channel and there is no goroutine that exists, the Go runtime will find it as deadlock and the program will panic
//   ----------------- NB ----------------- It is very important to understand that read and send operations in the channel are atomic.
// 	*/
// 	a := <-ch  //write 1o to the ch and store it in a
// 	fmt.Println(a)
// }

// func writeVal(ch chan int) {
// 	ch <- 10
// }

// error:
// func main() {
// 	ch := make(chan int, 0)
// 	a := <-ch
// 	fmt.Println(a)
// }

// another error:
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	ch := make(chan int)
// 	close(ch) // chant write in closed channel, only can read but it is always 0, secon return param is false when the channel is closed
// 	go writeVal(ch, &wg)
// 	wg.Wait()
// }

// func writeVal(ch chan int, wg *sync.WaitGroup)  {
// 	ch <-10
// 	wg.Done()
// }

// func main() {
// 	ch := make(chan int)
// 	close(ch)
// 	val, ok := <-ch // the second value for false is channel is closed
// 	if !ok {
// 		fmt.Printf("Value %v is returned. The channel is closed\n", val)
// 	}
// }

// import "fmt"

// func main() {
// 	ch := make(chan int)
// 	go func ()  {
// 		for i := 0; i < 10; i++ {
// 				ch <- i
// 		}
// 		close (ch)
// 	}()
// for {
// 	val, ok := <-ch
// 	if !ok {
// 		break
// 	}
// 	fmt.Println(val)
//} this code is equal to the one bellow:
// for val, ok := <-ch; ok; val, ok = <-ch {
// 	fmt.Println(val)
// } EQUAL TO THE CODE BELLOW:
// for val := range ch {
// 	fmt.Println(val) // best way because: we don't need a condition and when the channel is closed and all the values of the channel are read,
// 	//the range cause ends automatically:
// }
// }

/*
. In the unbuffered channel, the send operation will be blocked until one goroutine reads. The same applies for the read operation.
 The read operation blocks until one goroutine sends data to it.
*/
func main() {
	ch		:= make(chan int) // this line creates a new channel for sending ints between producer and consumer
	done	:= make(chan bool)
	go producer(ch, done) // go routines
	go consumer(ch)
	<-done
}

func producer(ch chan int, done chan bool) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("PRODUCER: sending %v\n", i)
	}
	close(ch)
	done <- true
}

// for val := range ch This loop continuously receives values from the ch channel until it is closed. Once closed, the loop terminates.
//  ----------- NB !! -------- ch <- value (send the value into the channel ch)
// ------------ NB !! -------- value := <-ch (receive a value from the channel ch and assign it to value)
// -----------  NB !! -------- <-done (wait to receive a value from the done channel)
func consumer(ch chan int) {
	for val:= range ch {
		fmt.Printf("CONSUMER: read %v\n", val)
	}
}
