package main

import (
	bs "binary-search/binary-search"
	"fmt"
)

func main() {

	res := bs.Search([]int{1}, 1)

	fmt.Print(res)
}
