package main

import "fmt"

func main() {
	fmt.Println("Hello, Welcome to my anything program on go language")
	foo()
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			//n,err := fmt.Println(i)
			n,_ := fmt.Println(i)
			fmt.Println(n)
			//fmt.Println(err)
		}
	}

	woo()
}

func foo() {
	fmt.Println("I am in foo function")
}

func woo() {
	fmt.Println("I am in woo and exiting")
}
