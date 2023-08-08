package main

/*
	#include <stdio.h>

	void helloWorld_c(char* message){
		printf("%s\n", message);
		return;
	}
*/
import "C"

func main() {
	C.helloWorld_c(C.CString("Hello World from C and GO"))
}
