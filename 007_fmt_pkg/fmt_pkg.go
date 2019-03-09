package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) // %T gives you the type
	fmt.Printf("%b\n", y) // %b formats an int as binary
	fmt.Printf("%X\n", y) // %X/x formats in hexadecimal (upper case/lower case)
	fmt.Printf("%#x\n", y) // %#X formats in hexadecimal with 0x before the number
	y = 911
	fmt.Printf("%#X\n", y)
	fmt.Printf("%#X\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#X\t%b\t%x\n", y, y, y)
	fmt.Println(s)
	fmt.Printf("%v", y)
}
