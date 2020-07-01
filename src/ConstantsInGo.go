package main

import "fmt"

func main() {

	//1. With constants, we can skip the types,
	const x = 3.14
	//fmt.Printf("%T\n", x)

	//2. when used in any operation, compiler will implicitly assume the type
	const stringConstant = "string constant"
	const integerConstant = 123
	//fmt.Printf("%v,%T\n", stringConstant + "_abc", stringConstant + "_abc")
	//fmt.Printf("%v,%T", integerConstant + 3, integerConstant + 3)

	/*-----------Enums in Go - here comes the 'iota'---------*/

	//3. we can declare the const block also - below will act like an Enum
	const (
		a = iota
		b = iota
		c = iota
	)

	//4. Another way to represent, we can skip iota for consecutive values;
	//5. Here, we can also change the initial value from 0 to 10
	const (
		a1 = iota + 10
		b1
		c1
	)
	fmt.Println(a,b,c)
	fmt.Println(a1,b1,c1)

	//6. We can also fix the first offset using underscore; acts as starting point ; first value is ignored
	//7. Using left shift operator in conjunction with the incrementing value of iota; cool way to declare memory size constants
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)


	fmt.Println(KB, MB, GB)

}
