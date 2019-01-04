package main

import "fmt"

func main() {
	type MyInt int
	type Ta *int
	type Tb *MyInt

	no1 := 10
	fmt.Println("after 'no1 := 10'")
	fmt.Println("no1=", no1)
	fmt.Println("&no1=", &no1)

	var no2 Ta = &no1
	fmt.Println("after 'var no2 Ta = &no1'")
	fmt.Println("no2=", no2)
	fmt.Println("*no2=", *no2)

	no3 := MyInt(30)
	fmt.Println("after 'no3 := MyInt(20)'")
	fmt.Println("no3=", no3)

	// no1 = no3	// not allowed
	// no3 = no1	// not allowed

	no4 := &no3
	fmt.Println("after 'no4 := &no3'")
	fmt.Println("*no4=", *no4)

	no5 := int(50)
	fmt.Println("after 'no5 := int(50)'")
	fmt.Println("no5=", no5)

	no6 := &no5
	fmt.Println("after 'no6 := &no5'")
	fmt.Println("no6=", no6)
	fmt.Println("*no6=", *no6)

	no4 = (*MyInt)(no6)
	fmt.Println("after 'no4 = (*MyInt)(no6)'")
	fmt.Println("no6=", no6)
	fmt.Println("no4=", no4)
	fmt.Println("*no6=", *no6)
	fmt.Println("*no4=", *no4)

	no4 = (*MyInt)((*int)(no2))
	fmt.Println("after 'no4 = (*MyInt)((*int)(no2))'")
	fmt.Println("no2=", no2)
	fmt.Println("no4=", no4)
	fmt.Println("*no2=", *no2)
	fmt.Println("*no4=", *no4)

	fmt.Println("value and types")
	fmt.Printf("no1=%d, no1.T=%T\n", no1, no1)
	fmt.Printf("no2=%d, no2.T=%T\n", no2, no2)
	fmt.Printf("no3=%d, no3.T=%T\n", no3, no3)
	fmt.Printf("no4=%d, no4.T=%T\n", no4, no4)
	fmt.Printf("no5=%d, no5.T=%T\n", no5, no5)
	fmt.Printf("no6=%d, no6.T=%T\n", no6, no6)
}
