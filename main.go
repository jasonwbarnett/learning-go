package main

import "fmt"
import "strconv"

func main() {
	name := "Jason Barnett"
	fmt.Printf("My name is: %s", name)
	var a [4]int
	fmt.Println(a)
	a[0] = 2
	fmt.Println(a)
	a[3] = 8
	fmt.Println(a)
	// s := make([]byte, 5)
	s := []byte{0,0,0,0}
	fmt.Println(len(s))
	fmt.Println(cap(s))

	var b string
	b = `⌘`;
	fmt.Println(b);
	fmt.Println(len("HELLO WORLD"))
	fmt.Println("Hello World"[2])
	fmt.Println(true && !false)
	fmt.Println((true && false) || (false && true) || !(false && false))

	var num int64
	var e error
	var oth sting
	num, e = strconv.ParseInt("12", 0, 0)
	fmt.Println(num)
	fmt.Println(e)
}
