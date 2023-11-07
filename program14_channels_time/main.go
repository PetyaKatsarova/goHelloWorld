package main

import (
	"fmt"
	"time"
)
/* 
The Ticker sends continuous events in the given time. We use Ticker when we need continuous events in a certain time.
We can read events from the channel of the ticker, i.e., C using <- operator:  msg := <-ticker.C 
The Timer is similar to the Ticker, but one difference: the Timer sends an event for only one time, whereas Ticker sends events continuously.
We cannot read a value from Timer more than once. If we do that, the program will panic.
We can use the range clause to read messages from the channel. When the channel is closed and all messages are read, the for loop ends.
*/

func main() {
	ticker1 := time.NewTicker(time.Millisecond*1000)
	ticker2 := time.NewTicker(time.Millisecond*1300)
	go receiver(ticker1, ticker2)
	time.Sleep(time.Millisecond * 3000) // wait 3 sec and stops ticker1
	ticker1.Stop()
	time.Sleep(time.Millisecond * 10000)
}

func receiver(ticker1 *time.Ticker, ticker2 *time.Ticker) {
	for {
		select { // reads from ticker1 or 2 respectevely
		case msg := <-ticker1.C:
			fmt.Println("Msg from ticker 1", msg)
		case msg := <-ticker2.C:
			fmt.Println("Msg from ticker 2", msg)
		}
	}
}
