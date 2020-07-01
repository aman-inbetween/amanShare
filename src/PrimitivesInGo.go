package main

import (
	"fmt"
)

//Scope: Global by conventions, Pascal casing(first letter Caps)
var GlobalVar = "Global variable"

//Scope: Local to package, camel casing(first letter Small)
var localVar = "Local variable package level"

//No compilation error for unused Variable defined at package level

func main() {
		
	//1. Variables are initialized to 0 by default - contrary to C++ where you get garbage(unless static)
	//2. Variables are scoped on the basis of blocks and also by conventions(Global in package scope if caps) - contrary to C, C++, JAVA
	//3. Implicit type conversion with loss of precision not allowed - contrary to C,C++
	//4. No Implicit conversions in arithmetic operations also(even int,int8,int32 are treated as different type for arithmatic)
	//5. Unused variables result in error instead of warning - contrary to C, C++, JAVA
	
	fmt.Println("-------------Understanding variables in Go---------------")

    //1. Different ways of declaration/assignment of variables
	var i int = 67

	var j float32
	j= 30.10

	k := 4 //type assumed at compile time

	fmt.Println(float32(i)+j)

	//2. to print the variable value and type both   - Similar to C
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

	
	//3. All arithmatic operations availble(+,-,*,/,%)
	//4. All bitwise operations availble(|, &, ^, ^&)
	//5. Bitshift operators available <<, >>


	//6. Below line with throw compile time error since arithmetic on different types
	//var xint16 int16 = 2;
	//var xint32 int32 = 3;
	//fmt.Println(xint16 + xint32)

	//7. Works after explicit type conversion
	//fmt.Println(xint16 + int16(xint32))


	//Go has also included native type support for complex numbers
	var n complex64 = 1+2i
	fmt.Printf("%v %T\n", n, n)
	
	
}
