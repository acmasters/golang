package main

/* %b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04
*/

import "fmt"

func main() {
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
}
