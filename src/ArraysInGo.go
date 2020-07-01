package main

import (
	"fmt"
)

func main() {

	//Contiguous memory allocation with each allocated block taking the space equal to given array type

	//1. declaring an array with predefined size
	var arr [3]int

	//2. Initializing array values
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Printf("\n %v, %T \n", arr, arr)

	//3. Another way declare and initialize array in one line
	var stringArr [3]string = [3]string{"apple", "mango"}
	fmt.Printf("\n %v  %T\n", stringArr, stringArr)

	//2. Another way to declare and initialize array in one line
	arr1 := [3]int{11, 12, 13}
	fmt.Printf("\n %v, %T \n", arr1, arr1)

	//3. Since we are already giving the literals, we can omit size and give [...] convention
	arr2 := [...]int{11, 12, 13}
	fmt.Printf("\n %v, %T \n", arr2, arr2)

	//4. length of the array
	fmt.Println(len(arr))

	//5. Below line will Out of bound exception - Similar to C, C++, JAVA etc
	//arr[4] = 4


	//6. Double dimension arrays - Matrix
	var matrix [2][2]int = [2][2]int{ [2]int{10,11}, [2]int{20,21}}
	fmt.Printf("\n %v, %T \n", matrix, matrix)

	/*--------------Copying an array--------------------*/
	//7. Copying an array results in creation of another array, it's not a reference to the original one
	var copyArr [3]int
	copyArr = arr    //this operation does another memory allocation and original array does not gets affected
	copyArr[2] = 300
	fmt.Println(arr, copyArr)

	/*------------Pointers and back-----------*/
	//8. In order to avoid the duplicate memory allocation, we can create references using '&' - Similar to C,C++
	var pointerToArr *[3]int
	pointerToArr = &arr
	pointerToArr[2] = 300
	fmt.Println(arr, pointerToArr)


	/*-----------Slices-------------*/
	//1. Slices are similar to arrays but are references instead of actual values
	//2. so when we create slice or a copy of slice, its by default a reference

	var sliceArr  = []int{200,300, 400, 500, 600}
	fmt.Printf("%v %T \n", sliceArr, sliceArr)
	fmt.Println(sliceArr, len(sliceArr), cap(sliceArr))

	var sliceArrCopy = sliceArr
	sliceArrCopy[1] = 3000
	//fmt.Println(sliceArr, len(sliceArr), cap(sliceArr))
	//fmt.Println(sliceArrCopy, len(sliceArrCopy), cap(sliceArrCopy))
	//e.g. use case: can be used for parallel batch processing of elements from a big array

	//3. Dynamically increasing the size of array using append() function
	//4. Go by default increases the capacity by ~twice of the initial slice - Similar to vectors in C++
	sliceArr = append(sliceArr, 700, 800)
	fmt.Println(sliceArr, len(sliceArr), cap(sliceArr))

}
