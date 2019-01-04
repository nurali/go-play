package main

import "fmt"

type emp struct {
	id   int
	name string
}

type empPtr *emp
type intPtr *int

// reference: https://go101.org/article/pointer.html
func main() {
	structPtrDemo()
	intPtrDemo()
}

func structPtrDemo() {
	fmt.Println()
	fmt.Println("** struct ptr demo **")
	e1 := new(emp)
	fmt.Println("after 'e1 := new(emp)'")
	fmt.Println("e1=", e1)
	fmt.Println("*e1=", *e1)

	e2 := &e1
	fmt.Println("after 'e2 := &e1'")
	fmt.Println("e2=", e2)
	fmt.Println("*e2=", *e2)
	fmt.Println("**e2=", **e2)

	fmt.Println()
	e3 := emp{}
	fmt.Println("after 'e3 := emp{}'")
	fmt.Println("e3=", e3)
	fmt.Println("&e3=", &e3)

	e4 := &e3
	fmt.Println("after 'e4 := &e3'")
	fmt.Println("e4=", e4)

	fmt.Println()
	e5 := &emp{}
	fmt.Println("e5 := &emp{}")
	fmt.Println("e5=", e5)
	fmt.Println("*e5=", *e5)

	fmt.Println()
	var e6 empPtr = &emp{}
	fmt.Println("after 'var e6 empPtr = &emp{}'")
	fmt.Println("e6=", e6)
	fmt.Println("*e6=", *e6)
	fmt.Println("&e6=", &e6)

	e7 := &e6
	fmt.Println("after 'e7 := &e6'")
	fmt.Println("e7=", e7)
	fmt.Println("*e7=", *e7)
	fmt.Println("**e7=", **e7)
}

func intPtrDemo() {
	fmt.Println()
	fmt.Println("** int ptr demo **")
	no1 := 10
	fmt.Println("after 'no1 := 10'")
	fmt.Println("no1=", no1)
	fmt.Println("&no1=", &no1)
	no2 := &no1
	fmt.Println("after 'no2 := &no1'")
	fmt.Println("no1=", no1)
	fmt.Println("&no1=", &no1)
	fmt.Println("no2=", no2)
	fmt.Println("*no2=", *no2)
	fmt.Println("&no2=", &no2)

	no1 = 11
	fmt.Println("after 'no1 = 11'")
	fmt.Println("no1=", no1)
	fmt.Println("no2=", no2)
	fmt.Println("*no2=", *no2)

	var no3 intPtr = &no1
	fmt.Println("after 'var no3 intPtr = &no1'")
	fmt.Println("no3=", no3)
	fmt.Println("*no3=", *no3)

	no1 = 12
	fmt.Println("after 'no1 = 12'")
	fmt.Println("no1=", no1)
	fmt.Println("*no2=", *no2)
	fmt.Println("*no3=", *no3)

	var no4 intPtr = no2
	fmt.Println("after 'var no4 intPtr = no2'")
	fmt.Println("no4=", no4)
	fmt.Println("*no4=", *no4)

	fmt.Println()
	no5 := new(int)
	fmt.Println("after 'no5 := new(int)'")
	fmt.Println("no5=", no5)
	fmt.Println("*no5=", *no5)
	*no5 = 10
	fmt.Println("after '*no5 = 10'")
	fmt.Println("no5=", no5)
	fmt.Println("*no5=", *no5)

	fmt.Println()
	no6 := int(10)
	fmt.Println("after 'no6 := int(10)'")
	fmt.Println("no6=", no6)
}
