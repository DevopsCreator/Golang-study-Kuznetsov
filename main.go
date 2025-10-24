package main

import "fmt"

func main() {

	var column int
	var line int

	for {
		fmt.Print("Enter column:")
		_, err := fmt.Scanf("%d", &column)
		if err != nil {
			fmt.Print("\nYou enter incorret value\nTry again!)\n")
			var discard string
			fmt.Scanln(&discard)
		} else{
			break
		}
	}
	for{
		fmt.Print("Enter line:")
		_, err := fmt.Scanf("%d", &line)
		if err != nil {
			fmt.Print("\nYou enter incorret value\nTry again!)\n")
			var discard string
			fmt.Scanln(&discard)
		} else{
			break
		}
	}

	for i := 0; i < line; i++ {
		if (i % 2) != 0 {
			for x := 0; x < column; x += 2 {
				fmt.Printf(" ")
				fmt.Printf("#")
			}
		} else {
			for x := 0; x < column; x += 2 {
				fmt.Printf("#")
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

}
