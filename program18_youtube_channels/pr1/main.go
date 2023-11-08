package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c := make(chan string)
// 	go count("sheep", c)

// 	for msg  := range c { // get everything sent to the channel
// 		// msg, open = <- c // here like above: store msg from c
// 		// if !open { break } // in range, no need to check manually if open, range looks after it
// 		fmt.Println(msg)
// 	}
// 	//!! receiver never closes a channel: he listens, if closed will panic
// }

// func count(thing string, c chan string) {
// 	for i := 0; i < 5; i++ {
// 		c <- thing
// 		time.Sleep(time.Millisecond * 500)
// 	}
// 	close(c) // if dont close the sending channel : deadlock: receiver keeps blocking
// }

// import (
// 	"fmt"
// )

//fatal error: all goroutines are asleep - deadlock!
// func main() {
// 	c := make(chan string, 2)
// 	c <- "hello" // send blocks till sth is ready to receive if no buffer number
// 	c <- "world"

// 	msg := <- c
// 	fmt.Println(msg)

// 	msg = <- c
// 	fmt.Println(msg)
// }

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func ()  {
// 		for {
// 			c1 <- "every 500ms"
// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}()

// 	go func ()  {
// 		for {
// 			c2 <- "every 2 sec"
// 			time.Sleep(time.Second * 2)
// 		}
// 	}()

// 	// always return at the same time, eventhough c1 is faster, cause we block ch
// 	// for {
// 	// 	fmt.Println(<- c1)
// 	// 	fmt.Println(<- c2)
// 	// }

// 	for {
// 		select { // doesnt block, allows to receive from whichever chan is ready
// 		case msg1 := <- c1:
// 		fmt.Println(msg1)
// 		case msg2 := <- c2:
// 			fmt.Println(msg2)
// 		}
// 	}
// }

import (
	"fmt"
)

// fibunacci

func main() {
// 	c := make(chan string, 2) // make a channel
// 	c <- "hello"           // send through the chan the string hello
// 	c <- "world"

// 	msg := <-c // msg receives from chan the string
// 	fmt.Println(msg)
// 	fmt.Println(<-c) // receive from channel c

	// workers code

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

// receive only to the jobs chan, send only to results chan
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

// quiet inneficient algorithm
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

/*
this will not execute cause send will block till sth is ready to receive! need to receive in separate go routine or buffer capacity, no need to receiving receiver
 c := make(chan string) // make a channel
 c <- "hello" // send through the chan the string hello

 msg := <-c // msg receives from chan the string
 fmt.Println(msg)
*/
