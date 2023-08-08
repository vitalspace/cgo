package main

/*
	#include <stdio.h>

	int factorial_c(int num) {
		if(num == 0 || num == 1) {
			return 1;
		} else {
			return num * factorial_c(num -1);
		}
	}

*/
import "C"
import "fmt"

func main() {
	num := 5
	result := C.factorial_c(C.int(num))
	fmt.Printf("The actorial od %d is %d", num, result)
}
