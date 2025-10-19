package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	Frist_Last(array)
}
func Zero_Based_indexs() {
	Zero_based := [5]int{1, 2, 3, 4, 5}
	fmt.Println(Zero_based)
}
func One_Based_indexs() {
	/*
		In Golang, arrays and slices are always 0-based
		indexed.

		That means the first element is at index 0, the
		second at 1, and so on. There is no built-in
		support for 1-based indexing like in some other
		languages (e.g., MATLAB, Lua).
	*/
}
func Static_and_Dynamic_arrays() {
	static_array := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(static_array)

	Dynamic_array := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Dynamic_array)
}

func For_loop_traversal() {
	array := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}
func While_loop_traversal() {
	array := []int{1, 2, 3, 4, 5, 6}
	n := 0
	for n < len(array) {
		fmt.Println(n)
		n++
	}
}
func Range_loop_traversal() {
	array := []int{1, 2, 3, 4, 5, 6}
	for key, value := range array {
		fmt.Printf("key:%v value:%v", key, value)
	}
}

func Forward_traversal() {
	array := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i])
		fmt.Println("->")
	}
}

/*
1,2,3,4,5
0,1,2,3,4
*/
func Reverse_traversal(array [5]int) {
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Println(array[i])
	}
}
func Reading_Input_from_console() {
	var n int
	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&n)

	var arr = [100]int{} // fixed-size array (max 100 elements)

	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Array elements are:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}

func Reading_input_from_file() {
	bytesdata, err := os.ReadFile("C:/Users/gajjala mahindra/OneDrive/Desktop/Important_Coding_Questions/array/TRAVERSAL/input.txt")

	if err != nil {
		fmt.Println(err)
	}
	slice_strings := strings.Fields(string(bytesdata))
	fmt.Println(slice_strings)
}
func Frist_Last(array [5]int) {
	//Finding first/last index of a value
	fmt.Println("Finding Frist index of a value", array[0])
	fmt.Println("Finding Last index of a value", array[len(array)-1])
}
