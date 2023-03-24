package main

import (
	"fmt"
	sip "search-insert-position/search-insert-position"
)

func main() {
	res := sip.SearchInsert([]int{1, 2, 3}, 3)

	fmt.Print(res)
}
