package main

import "fmt"

func main() {
	type MyInt int
	type Ta *int
	type Tb *MyInt

	var no1 Ta
	var no2 *int
	var no3 Tb
	var no4 *MyInt
	var no5 int
	fmt.Println("no5=", no5)

	_ = no1 == no2
	_ = no3 == no4
	// _ = no1 == no3
	// _ = no1 == no4
	// _ = no2 == no3
	// _ = no2 == no4

	_ = no1 == nil
	_ = no2 == nil
	_ = no3 == nil
	_ = no4 == nil
	// _ = no5 == nil
}
