# goHelloWorld
Learning points:
1. in your program (at the root), from the console you need run:
              go mod init main 
// because main is the package name i chose: this way all files with the same declared package name are connected
// in all my files in program, on row 1, i declare "package main"

2. To run the program, in the console, at the root of the program u need to add all files used in the main method, example:
            go run main.go functions.go manyReturnVals.go

3. Pass by value: