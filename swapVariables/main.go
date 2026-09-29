package main

import "fmt"

func main() {

    // c 3 
	a1 := 2143
	a2 := 33
	fmt.Printf("a1:%v a2:%v\n", a1, a2)	
	var temp int
	temp = a1 
	a1 = a2 
	a2 = temp
	fmt.Printf("a1:%v a2:%v\n", a1, a2)	
	// без 3 
	a3 := 123 
	a4 := 321 
	fmt.Printf("a3:%v a4:%v\n", a3, a4)	
	a3, a4 = a4, a3 
	fmt.Printf("a3:%v a4:%v\n", a3, a4)	

}