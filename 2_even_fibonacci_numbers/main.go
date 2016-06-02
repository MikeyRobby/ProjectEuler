package main

import (
	"fmt"
)



//Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
//
//1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
//By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

func fibsum(max int) int{
	var totalSum int
	for i, j := 0, 1; j < max ; i, j = i+j,i  {
		if i % 2 == 0 {
			totalSum += i
		}
	}
	return totalSum

}

func main() {
	fmt.Print(fibsum(4000000))
}
