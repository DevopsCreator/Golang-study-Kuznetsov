package main

import "fmt"

func main() {

	column := 8
	line := 8

	for	i := 0;	i < line;	i++ {
		for i := 0; i < column; i++ {
			fmt.Printf("# ")
		}
		fmt.Println("\n")
	}


}
