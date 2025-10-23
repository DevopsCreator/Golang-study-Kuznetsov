package main

import "fmt"

func main() {

	column := 8
	line := 8

	for	i := 0;	i < line;	i++ {
		if (i % 2) != 0 {
			for x := 0; x < column; x+=2 {
				fmt.Printf(" ")
				fmt.Printf("#")
			}
		} else {
			for x := 0; x < column; x+=2 {
				fmt.Printf("#")
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}


}
