package main

import "fmt"



//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
//Find the sum of all the multiples of 3 or 5 below 1000.

func math(number int)int {
	var multhree, mulfive int
	for i := 0; i < number; i++ {
		if i % 5 == 0 {
			mulfive += i
		}else if i % 3 == 0 {
			multhree += i
		}
	}
	total := mulfive + multhree
	return total
}

func main() {
	fmt.Print(math(1000))

}
