package main

import "fmt"

func main() {
	Check_key_exists_in_a_map(3)
}
func Check_key_exists_in_a_map(keydata int) {
	//Check if a key exists in a map.
	mapdata := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	for key, value := range mapdata {
		fmt.Println(key, value)
		if key == keydata {
			fmt.Println("exist:", value)
		}
	}
}
