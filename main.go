package main

import (
	"fmt"

	"math"
)


func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int)int {
	return num1 + num2
}

//Hello world
func main() {
	fmt.Println("Hello, World!")


//Data types
// var name = "Israel"
// fmt.Println(name)

// var age = 37
// fmt.Println(age) 

// const isCool = true

//Shorthand
// size := 1.3

name, age, email, size, isCool := "Israel", 37, "Israel@com", 1.3, true
fmt.Printf("%T %T %T %T %T\n", name, age, isCool, size, email)

//Packages
fmt.Println(math.Floor(2.7))
fmt.Println(math.Ceil(2.7))
fmt.Println(math.Sqrt(16))
fmt.Println(math.Exp(2))


//Functions

fmt.Println(greeting("Israel"))
fmt.Println(getSum(3, 4))

// //Arrays
//  var fruitArr [2]string

//  //Assign values
//  fruitArr[0] = "Apple"
//  fruitArr[1] = "Orange"
 

 //Declare and assign
// fruitArr := [2]string{"Apple", "Orange"}

//  fmt.Println(fruitArr)
//  fmt.Println(fruitArr[0])

fruitArr := []string{"Apple", "Orange"}
fmt.Println(fruitArr)

 //Slices
 fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
 fmt.Println(fruitSlice)

numArr := []int{1, 2, 3, 4, 5}
fmt.Println(numArr)

//Conditionals
x := 15
y := 10

//if else
if x <= y {
	fmt.Printf("%d is less than %d\n", x, y)
} else {
	fmt.Printf("%d is less than %d\n", y, x)
}

//else if
color := "yellow"
if color == "blue" {
	fmt.Println("Color is blue")
}else if color == "green" {
	fmt.Println("Color is green")
} else {
	fmt.Println("Color is not blue or green")
}

//Switch
switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	case "green":
		fmt.Println("Color is green")
	default:
		fmt.Println("Color is not red, blue or green")
}

//Loops
// i := 1
// for i <= 10 {
// 	fmt.Println(i)
// 	i++

//Short method
// for i := 0; i < 5; i++ {
// 	fmt.Printf("Number %d\n", i)
// }

// for i := 1; i < 21; i++ {
// 	if i % 3 == 0 && i % 5 == 0 {
// 		fmt.Printf("FizzBuzz ")
// 	} else if i % 3 == 0 {
// 		fmt.Printf("Fizz ")
// 	} else if i % 5 == 0 {
// 		fmt.Printf("Buzz ")
// 	}
// 	fmt.Println(i)
// }

//Maps
emails := make(map[string]string)
emails["Bob"] = "bob@gmail.com"
emails["Sharon"] = "sharon@gmail.com"
emails["Mike"] = "Mike@gmail.com"

fmt.Println(len(emails["Sharon"]))

cars := make(map[string]string)

cars["Brand"] = "Toyota"
cars["Model"] = "Camry"
cars["Year"] = "1964"

fmt.Println(cars["Brand"],cars["Model"],cars["Year"])
delete(emails, "Sharon")
fmt.Println(emails)




}




