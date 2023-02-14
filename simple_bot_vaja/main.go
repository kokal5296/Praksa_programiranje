package main

import (
	"fmt"
	"strconv"
)

func main() {
	bot_name := "Koki"
	birth_year := 2023
	u_name := ""

	fmt.Println("Hello! My name is ", bot_name)
	fmt.Println("I was created in ", birth_year)

	fmt.Println("Please, remind me your name.")
	fmt.Print("> ")
	fmt.Scan(&u_name)
	fmt.Println("What a great name you have, " + u_name + "!")

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")
	var rem3, rem5, rem7 int
	fmt.Print("> ")
	fmt.Scan(&rem3)
	fmt.Print("> ")
	fmt.Scan(&rem5)
	fmt.Print("> ")
	fmt.Scan(&rem7)
	age := (rem3*70 + rem5*21 + rem7*15) % 105
	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programing!")

	fmt.Println("Now I will prove to you that I can count to any number you want.")
	var u_num, num1 int
	fmt.Scan(&u_num)
	for num1 = 0; num1 <= u_num; num1++ {
		fmt.Println(num1, "!")
	}
	fmt.Println("Completed, have a nice day!")
	var odg int

	fmt.Println("Let's test your programming knowledge.")
	fmt.Println("Why do we use methods?") //
	fmt.Println("1. To repeat a statement multiple times.")
	fmt.Println("2. To decompose a program into several small subroutines.")
	fmt.Println("3. To determine the execution time of a program.")
	fmt.Println("4. To interrupt the execution of a program.")
	for odg != 2 {
		fmt.Scan(&odg)
		if odg == 2 {
			fmt.Println("Congratulations, have a nice day!")
		} else {
			fmt.Println("Please, try again.")
		}
	}
}
