package main

// import ( // packages from the standard library
// 	"fmt"
// 	"sort"
// 	"strings"
// )

// func main() {

// 	ages := []int{45, 20, 35, 30, 75, 60, 50, 35} // this is a slice, if i put number inside [3] it gets arrray with fixed length, not changable
// 	sort.Ints(ages) // changes original slice
// 	fmt.Println(ages)

// 	index := sort.SearchInts(ages, 30) // from sorted ages search index of 30
// 	fmt.Println("index in the sorted slice:", index)
// 	fmt.Println("not existance int index in the sorted slice:",sort.SearchInts(ages, 100))

// 	greeting := "hello there friends!"
// 	fmt.Println(strings.Contains(greeting, "hello!"))
// 	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) // doesnt change original str, only returns a new one
// 	fmt.Println(strings.ToUpper(greeting)) 
// 	fmt.Println(strings.Index(greeting, "ll")) //2
// 	fmt.Println(strings.Split(greeting, "e")) // returns an array [h llo th r  fri nds!]
// }