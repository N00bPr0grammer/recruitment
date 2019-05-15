package sort

import (
	"fmt"
	"sort"
)

//barcharts
func Barcharts() {
	var x = "|"
	y := []int{1,4,5,6,8,2}
	for index, x := range y {
		fmt.Println("%x")
	}
	return
}

//ascending sorting
func Sort(){
	var x = "|"
	y := []int{1,4,5,6,8,2}
	for index, x := range y {
		fmt.Println("%x")
	}
	return sort(x)
}

//reverse sorting
func Sort2() {
	var x = "|"
	y := []int{1,4,5,6,8,2}
	for index, x := range y {
		fmt.Println("%x")
	}
	return sort.Reverse(y)
}
