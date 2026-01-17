package main

import "fmt"

func main() {

	// switch case
	// switch value
	// case matcher do statement
	// case matcher do statement
	// var day string = "Wednesday"
	// switch day {
	// case "Monday":
	// 	fmt.Println("Thu 2")
	// case "Tuesday":
	// 	fmt.Println("Thu 3")
	// case "Wednesday":
	// 	fmt.Println("Thu 4")
	// case "Thurday":
	// 	fmt.Println("Thu 5")
	// case "Friday":
	// 	fmt.Println("Thu 6")
	// case "Saturday":
	// 	fmt.Println("Thu 7")
	// case "Sunday":
	// 	fmt.Println("Chu Nhat")
	// }
	// var day string = "Monday"
	// switch {
	// case day == "Monday":
	// 	fmt.Println("Thu 2")
	// case day == "Tuesday":
	// 	fmt.Println("Thu 3")
	// case day == "Wednesday":
	// 	fmt.Println("Thu 4")
	// case day == "Thurday":
	// 	fmt.Println("Thu 5")
	// case day == "Friday":
	// 	fmt.Println("Thu 6")
	// case day == "Saturday":
	// 	fmt.Println("Thu 7")
	// case day == "Sunday":
	// 	fmt.Println("Chu Nhat")
	// }

	//not negative or not postive
	var number = 0
	switch {
	case number >= 0:
		fmt.Println("not Negative")
		fallthrough
	case number <= 0:
		fmt.Println("Not Positive")
	}

}
