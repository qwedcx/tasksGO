package main

import "fmt"

func main() {
	var a int
	var b float64
	var c string
	var d bool

	var aa = 1
	var bb = 1.1
	var cc = "jkhk"
	var dd = true

	aaa := 2143
	bbb := 1.2
	ccc := "asdasdasdasd"
	ddd := false

	fmt.Printf("%v = %T\n", a, a)
	fmt.Printf("%v = %T\n", b, b)
	fmt.Printf("%v = %T\n", c, c)
	fmt.Printf("%v = %T\n", d, d)

	fmt.Printf("%v = %T\n", aa, aa)
	fmt.Printf("%v = %T\n", bb, bb)
	fmt.Printf("%v = %T\n", cc, cc)
	fmt.Printf("%v = %T\n", dd, dd)

	fmt.Printf("%v = %T\n", aaa, aaa)
	fmt.Printf("%v = %T\n", bbb, bbb)
	fmt.Printf("%v = %T\n", ccc, ccc)
	fmt.Printf("%v = %T\n", ddd, ddd)
}