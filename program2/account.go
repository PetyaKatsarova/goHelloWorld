package main

import (
	"fmt"
	"sync"
)

type Account struct {
	balance int
	mutex sync.Mutex
}
/*
  sync.Mutex is a synchronization primitive that can be used to ensure that only one goroutine can access a particular section of code at a time. 
 This is used to prevent race conditions, where the outcome of the program execution may change due to the order of operations performed by multiple
  goroutines concurrently, which can lead to unpredictable behavior and bugs.The term "mutex" stands for "mutual exclusion," and a sync.Mutex provides 
 two primary methods: Lock() Unlock()

Lock():

When a goroutine reaches a point where it must access a shared resource (like a shared variable), it calls Lock() on the mutex associated with that resource.
If no other goroutine holds the lock, the Lock() call will succeed, and the calling goroutine will own the mutex. This means it can safely access the shared resource.
If another goroutine already holds the lock, the calling goroutine is blocked until the mutex is unlocked.
Unlock():

Once the goroutine that holds the mutex has finished working with the shared resource, it must call Unlock() to release the mutex.
Releasing the mutex unlocks it, and any other goroutines waiting for the lock can proceed to try to lock it themselves.
It's important to ensure that Unlock() is called in such a way that it will be executed even if an error occurs during the execution of the locked section of code, to prevent deadlocks. This is typically done using a defer statement directly after the mutex is locked, like so:

go
Copy code
var mu sync.Mutex

func criticalSection() {
    mu.Lock()
    defer mu.Unlock() // This ensures that the mutex is unlocked when the function returns.

    // ... critical section with shared resource ...
}
*/
