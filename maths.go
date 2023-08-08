package main

/*
	#include <stdio.h>

	int add_c(int numa, int numb) {
		return numa + numb;
	}

	int sub_c(int numa, int numb) {
		return numa - numb;
	}

	int mult_c(int numa, int numb) {
		return numa * numb;
	}

	int div_c(int numa, int numb) {
		return numa / numb;
	}

	int rest_c(int numa, int numb) {
		return numa % numb;
	}
*/
import "C"
import "fmt"

func main() {

	numa := 5
	numb := 2

	fmt.Printf("%d + %d = %d \n", numa, numb, C.add_c(C.int(numa), C.int(numb)))
	fmt.Printf("%d - %d = %d \n", numa, numb, C.sub_c(C.int(numa), C.int(numb)))
	fmt.Printf("%d * %d = %d \n", numa, numb, C.mult_c(C.int(numa), C.int(numb)))
	fmt.Printf("%d / %d = %d \n", numa, numb, C.div_c(C.int(numa), C.int(numb)))
	fmt.Printf("%d rest %d = %d \n", numa, numb, C.rest_c(C.int(numa), C.int(numb)))

}
