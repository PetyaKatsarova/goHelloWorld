package main

// import (
// 	"fmt"
// 	"time"
// 	"sync"
// )

// func main() {
// f1 is exitted half way after main method is done
// 	go f1("F1") // like async in js
// 	go f1("F2")
// 	fmt.Println("Sleeping for 5 sec")
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Main completed")
// }

// func f1(name string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%v: i %d\n", name, i)
// 		time.Sleep(1 * time.Second)
// 	}

// example 2: --------------- WAIT GROUP
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go f1("F1", &wg)
// 	go f1("F2", &wg)
// 	fmt.Println("Main: waiting for gorouitnes to fihish")
// 	wg.Wait()
// 	fmt.Println("Main completed")

// func f1(name string, wg *sync.WaitGroup)  {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%v: i %d\n", name, i)
// 		time.Sleep(1 * time.Second)
// 	}
// 	wg.Done()
// }

// ----------------- EXAMPLE 3: LOCK CRITICAL SECTION TO AVOID RACE CONDITION ------------------------------
// A drawback of the Mutex lock: Since it locks for other goroutines to access the critical section, it slows down the overall execution
//  of the program.
// var amount = 1000
// func main() {
// 	var wg sync.WaitGroup
// 	var m  sync.Mutex
// 	wg.Add(100)
// 	for i := 0; i < 100; i++ {
// 		go withdraw(&wg, &m)
// 	}
// 	wg.Wait()
// 	fmt.Println(amount)
//  }

//  func withdraw(wg *sync.WaitGroup, m *sync.Mutex)  {
// 	defer wg.Done() // always release the func from the wg, even if returns error
// 	m.Lock()
// 	amount = amount - 1
// 	m.Unlock()
//  }

// ------------------- GOMAXPROCS ----------------------------------
//We can use the Goexit function to stop the execution of the goroutine. The Goexit function will stop a goroutine in which it is called

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2) // will run paralel: if 2 cores; will run on both
	var wg sync.WaitGroup
	wg.Add(2)
	go f1("F1", &wg)
	go f1("F2", &wg)
	fmt.Println("waiting for goroutines to finish")
	wg.Wait() // dont complete main untill all other threads are completed
	fmt.Println("main completed")
}

func f1(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		if i == 5 {
			runtime.Goexit() // similar to return: exits the process
		}
		fmt.Printf("%v: i %d\n", name, i)
	}
	wg.Done()
}

// defer wg.Done() explained **********************************************
/*
	defer wg.Done() schedules the wg.Done() call to be executed just before the withdraw function returns. The wg.Done() method is
	 part of the sync.WaitGroup type and is used to signal that a goroutine has finished its work. When you add a count with wg.Add(n),
	  you're telling the WaitGroup that n pieces of work have started. By deferring wg.Done(), you ensure that:

No matter how the withdraw function exits (normal completion or early return due to an error condition), wg.Done() will be called.
This prevents a scenario where a goroutine might exit without signaling the WaitGroup, which would cause anything waiting on wg.Wait()
 to hang indefinitely.The wg.Done() call is clearly visible at the start of the function, which improves readability.
  It's easy to see that this function will signal the WaitGroup when it's done, without having to read through the entire
   function body. In more complex functions, if there are multiple return statements or the possibility of a panic, using defer
    ensures the WaitGroup is always decremented, avoiding deadlocks or leaked goroutines.In simple words, defer wg.Done()
	 in your withdraw function ensures that wg.Done() is called right before the function exits, no matter how the function exits.
	 This is important for the WaitGroup to work correctly, as it will only unblock wg.Wait() when all Done() calls have been
	  made corresponding to the number of Add() calls.
*/
