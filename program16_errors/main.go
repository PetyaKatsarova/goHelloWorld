package main

/*
In Go, a map is a built-in data type that associates keys with values. The keys in a map are unique: there cannot be duplicate keys, although you can have duplicate values. Maps are also known as associative arrays,
dictionaries, or hash tables in other programming languages.
*/

import "fmt"

// func main() {
// 	id, err := create("John", -7)
// 	if err != nil {
// 		fmt.Println("Error while creating student")
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Student created with id =", id)
// }

// type student struct {
// 	id int
// 	name string
// 	age int
// }

// var stuMap = make(map[int]student)

// func create(name string, age int) (int, error) {
// 	if name == "" {
// 		return 0, fmt.Errorf("Invalid Name, Name cant be empty")
// 	} else if age <= 0 {
// 		// return 0, fmt.Errorf("Invalid age %v, age cant be <= 0", age)
// 		return 0, InvalidAgeError{age}
// 	}
// 	id := len(stuMap) + 1;
// 	stuMap[id] = student{id, name, age}
// 	return id, nil
// }

// type InvalidAgeError struct {
// 	age int
// }

// func (i InvalidAgeError) Error() string {
// 	return fmt.Sprintf("Invalid age %v, age cant be <= 0", i.age)
// }
/*
---- PANIC ------------------------------------------
The unexpected runtime error stops the execution flow of the program abruptly. The most common examples of panic are divide by zero, index out of range, etc. When panic is encountered, the program terminates,
 deferred functions are called, and control goes to the caller. This process continues till the control reaches the first method of the goroutine call stack.Finally, it prints the stack trace on the console:
*/

// func main() {
// 	fmt.Println("Dividing by zero")
// 	a := 10
// 	b := 0
// 	c := a / b
// 	fmt.Println("Result:", c)
// }

/*
-------------- RECOVER ---------------------------------
The built-in function recover recovers from panic. The program usually executes after the recover call. It should be called from the defer function; otherwise, it will not have any impact.

The recover function can be called from any function in the call stack. It is not necessary to call the recover function the same function from where panic is encountered.
In the code given below, the recover function is called in the fn2 function. After the fn2 call, the program doesn't terminate. Instead, it regains the flow and ends the program gracefully.
It is important to note that the last line of fn2 doesn't execute: The recover function returns nil if there is no panic in the call stack
if the recover function is called without the defer function, it will not have any impact and it will always return nil
*/

// func main() {
// 	fmt.Println("Start of Main Func")
// 	defer func() { fmt.Println("Defer in Main func") }()
// 	fn1()
// 	fmt.Println("End of Main Func")
// }
// func fn1() {
// 	fmt.Println("Start of Fn1 Func")
// 	fmt.Println("Recovered", recover()) // need KEY WORD DEFER for the recover to panic!!
// 	defer func() {
// 		fmt.Println("Defer in Fn1 func")
// 	}()
// 	panic("Panicing from fn1")
// }

// func main() {
// 	fmt.Println("Start of main func")
// 	defer func() {
// 		fmt.Println("Dever in main func")
// 	}()
// 	fn1()
// 	fmt.Println("end of main func")
// }

// func fn1() {
// 	fmt.Println("Start of fn1 func")
// 	defer func() { fmt.Println("Defer in fn1 func") }()
// 	fn2()
// 	fmt.Println("end of fn1 func")
// }

// func fn2() {
// 	fmt.Println("Start of Fn2 Func")
// 	defer func() {
// 		fmt.Println("Defer in fn2")
// 		if rec := recover(); rec != nil {
// 			fmt.Println("Recoverd:", rec)
// 		}
// 	}()
// 	fn3()
// 	fmt.Println("End of fn2")
// }

// func fn3() {
// 	fmt.Println("Start of Fn3 Func")
// 	defer func() { fmt.Println("Defer in fn3 func") }()
// 	panic("calling panic explicitly in fn3")
// }

/*
---------------------------------- DEFER---------------------------------------------------------

In Go, defer is a keyword that is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup. defer is often used where e.g., ensure or finally would be used in
other languages.When a defer statement is encountered, the function call following it is not executed immediately. Instead, it's scheduled to be called after the function that contains the defer statement has
 finished executing, just before it returns. The deferred calls are executed in last-in, first-out order, which means that if you have multiple defer statements, the last one in the code is executed first.
Defer is commonly used to close files or network connections, unlock mutexes, or perform some kind of finalization step, regardless of how the function exits (whether it's a normal return or by panicking).

func main() {
    f, err := os.Open("filename.txt")
    if err != nil {
        log.Fatal(err)
    }
    // This line will ensure that the file is closed
    // after the surrounding function's scope is done.
    defer f.Close()

    // ... do something with the file ...

    // At the end of the function, f.Close() will be called.
}
In this example, defer f.Close() ensures that the Close method on f (which is a file handle) is called when the function main finishes, regardless of whether it finishes successfully or with an error.
Deferred functions are executed after any return statements, but before the function actually returns to the calling function.
Arguments to deferred functions are evaluated when the defer statement is executed, not when the actual deferred call is executed.
Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in, first-out order.
Deferred functions are executed even if a runtime panic occurs in the function, which makes them useful for releasing resources that would otherwise be leaked during panic recovery.
Using defer can make your code cleaner and much easier to understand, especially when dealing with resource management and error handling.
*/

/*

In Go, the standard way to create a new error object is by using the New function from the errors package. So the correct answer is:

New()

Here's an example of how it's used:

go
Copy code
import (
    "errors"
    "fmt"
)

func main() {
    err := errors.New("this is a new error")
    fmt.Println(err)
}
*/

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from error: ", err)
		}
	}()

	arr := [5]int{1,2,3,4,5}
	i := 5
	fmt.Println(arr[i])
}
