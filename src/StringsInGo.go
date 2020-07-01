package main

import (
	"fmt"
)

func main() {

	var str string = "wonder land!"
	fmt.Printf("%v, %T \n", str, str)

	//1. Strings in Go are immutable -  Similar to JAVA
	str = "another Value for same string"
	fmt.Println(str)

	//2. Accessing using [] returns a byte
	fmt.Printf("%v, %T \n", str[0], str[0])

	//3.  Below code with throw compile time error since arithmetic on different types -- Contrary to JAVA
	i := 67
	//fmt.Println("some string_" + i)

	//4.  Below line with give unexpected result  -- Similar to C, C++
	fmt.Println("some string_" + string(i))


	/*----------------Rune------------------*

	*/




}
