/*
Given a binary number, write a program that prints 1
if given binary number is a multiple of 3.  Else prints 0.
The given number can be big upto 2^100.
It is recommended to finish the task using one traversal of input binary string.
*/

package misc

import "fmt"

/*
"101"
*/

func shift(n int) {
	fmt.Println(n >> 1)
}

// func Run() {
// 	shift(10)
// }
