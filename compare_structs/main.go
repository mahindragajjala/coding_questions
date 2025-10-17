package main

import "fmt"

type struct1 struct {
	value1 int
	value2 int
}
type struct2 struct {
	value1 int
	value2 int
}

func main() {
	s1 := struct1{
		value1: 1,
		value2: 2,
	}
	s2 := struct2{
		value1: 1,
		value2: 2,
	}
	fmt.Println(s1, s2)
	/* 	if s2 == s1{

	   	} */
}
