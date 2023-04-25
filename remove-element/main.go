package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	if position == "left" {
		return arr[1:]
	} else if position == "right" {
		return arr[:len(arr)-1]
	} else {
		return arr
	}
}

func main() {
	arr := []any{1, 2, 3, 4, 5}
	leftArr := removeLeftRight(arr, "left")
	fmt.Println(leftArr)
	rightArr := removeLeftRight(arr, "right")
	fmt.Println(rightArr)
}
