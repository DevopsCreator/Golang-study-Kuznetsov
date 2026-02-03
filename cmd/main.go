package main

import (
	"bufio"
	"fmt"
	"os"
)

func getNameOfUsers(firstUser, secondUser *string) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter first Username:")
	scanner.Scan()
	*firstUser = scanner.Text()
	fmt.Print("Enter second Username:")
	scanner.Scan()
	*secondUser = scanner.Text()
}

func getColumnAndLine(column, line *int) {

	for {
		fmt.Print("Enter column:")
		_, err := fmt.Scanf("%d", column)
		if err != nil {
			fmt.Print("\nYou enter incorret value\nTry again!)\n")
			var discard string
			fmt.Scanln(&discard)
		} else if *column > 26 {
			fmt.Print("\nYou enter incorret value column must be less 26\n")
		} else {
			break
		}
	}
	for {
		fmt.Print("Enter line:")
		_, err := fmt.Scanf("%d", line)
		if err != nil {
			fmt.Print("\nYou enter incorret value\nTry again!)\n")
			var discard string
			fmt.Scanln(&discard)
		} else if *line < 5 {
			fmt.Print("\nYou enter incorret value line must be more 5\n")
		} else {
			break
		}
	}

}

func alphabetPrint(alphabetStart, column int) {

	for x := alphabetStart; x < column+alphabetStart; x++ {
		if x == alphabetStart{
			fmt.Print("  ")
		}
		if x == alphabetStart && column >= 10{
			fmt.Print(" ")
		}
		fmt.Printf("%c ", x)
	}
	fmt.Println()
}

func main() {

	var column int
	var line int
	var firstUser string
	var secondUser string
	var alphabetStart int = 65
	blackFirstLine := [26]rune{9820, 9822, 9821, 9819, 9818, 9821, 9822, 9820, 9820, 9822, 9821, 9819, 9818, 9821, 9822, 9820, 9820, 9822, 9821, 9819, 9818, 9821, 9822, 9820, 9820, 9822}
	whiteFirstLine := [26]rune{9814, 9816, 9815, 9813, 9812, 9815, 9816, 9814,9814, 9816, 9815, 9813, 9812, 9815, 9816, 9814,9814, 9816, 9815, 9813, 9812, 9815, 9816, 9814,9814, 9816}
	getNameOfUsers(&firstUser, &secondUser)

	getColumnAndLine(&column, &line)

	// Print Letters for chessdesk
	alphabetPrint(alphabetStart, column)

	for i := line; i > 0; i-- {
		fmt.Print(i)
		if i == line {
			for x := 0; x < column; x++ {
				fmt.Printf(" %c", blackFirstLine[x])
			}
			fmt.Println("  Player One:",firstUser)
		} else if i == line-1{
			for x := 0; x < column; x++ {
				fmt.Printf(" %c",9823)
			}
			fmt.Println("  Player Two:",secondUser)
		} else if i == 1{
			for x := 0; x < column; x++ {
				fmt.Printf(" %c", whiteFirstLine[x])
			}
		}else if i == 2{
			for x := 0; x < column; x++ {
				fmt.Printf(" %c",9817)
			}
			fmt.Println()
		}else{
			if i % 2 != 0 {
				fmt.Printf(" ")
				for x := 0; x < column; x++ {
					if x % 2 == 0{
						fmt.Printf("%c",11035)
					}else{
						fmt.Printf("%c",11036)
					}
				}
			}else{
				fmt.Printf(" ")
				for x := 0; x < column; x++ {
					if x % 2 == 0{
						fmt.Printf("%c",11036)
					}else{
						fmt.Printf("%c",11035)
					}
				}
			}
			fmt.Println()
		}
	}

}
