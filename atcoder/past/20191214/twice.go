package main

// https://atcoder.jp/contests/past201912-open/tasks/past201912_a

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scan(&input)
	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("error")
		return
	}
	if i > 999 {
		fmt.Println("error")
		return
	}
	fmt.Println(i * 2)
}
