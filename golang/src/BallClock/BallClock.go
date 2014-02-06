package main

import (
	"fmt"
)

func main() {
	var i int

	for  {
		_, err := fmt.Scanf("%d\n", &i)
		
		if err != nil {  fmt.Println("ERROR: ", err); break }
		if i == 0 { break }
		
		if i >= 27 && i <= 127 {
			c := NewClock(i)
			if cycles, e := c.getLCM(); e == nil {
				fmt.Println(i, "balls cycle after ", cycles / 2, " days")
			} else {
				fmt.Println("ERROR: %v", e)
			}
		}
	}
}
