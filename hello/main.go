package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
	"unicode/utf8"
)

func foo() string {
	return "yolo"
}

func printMe(param string) {
	fmt.Println(param)
}

func main() {
	fmt.Println("HELLO FROM GO")

	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	//Arthimetic operations in go
	var floatNum1 float32 = 10.1
	var intNum32 int32 = 2
	var convertedInt float32 = float32(intNum32)
	var result float32 = floatNum1 + convertedInt
	fmt.Println("This is the result after converting int into float", result)

	//Division in go doesnt give a floating point number
	var intNum1 float32 = 5
	var intNum2 float32 = 3
	fmt.Println(intNum1 / intNum2)                            //gives the quotient
	fmt.Println(math.Mod(float64(intNum1), float64(intNum2))) //gives the remainder

	//Strings
	var myString string = "Hello world"
	var secondLiteral string = "This" + " " + "great"

	var resultString string = myString + " " + secondLiteral
	fmt.Println(resultString)
	fmt.Println(len("你好"))                    // This is important because strings in Go are UTF-8 encoded, and some characters may consist of multiple bytes.
	fmt.Println(utf8.RuneCountInString("你好")) // This gives the correct count of Unicode characters (runes), regardless of how many bytes each character takes.

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println((myBoolean))

	var intNum3 string
	fmt.Println(intNum3) //defaults to an empty string if the variable type is a string or 0 if it is anything else

	var myVar string = foo()
	fmt.Println(myVar)

	var1, var2 := 1, 2 //The := syntax in Go is used for short variable declaration, where Go automatically infers the type of the variable based on the assigned value, and it's typically used within functions.
	fmt.Println(var1, var2)

	const myConst string = "const value"
	const PI float32 = 3.14 //immutable down the execution line
	fmt.Println(PI, myConst)

	var printMeParam string = "this is you"
	printMe(printMeParam)

	var numerator int = 10
	var denominator int = 2

	var resultDivision, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Printf("The result of the integer division is %v with remainder %v", resultDivision, remainder)
	// }

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("There is result of division is %v\n", resultDivision)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v\n", resultDivision, remainder)
	}

	//ARRAY
	// var intArr [3]int32 = [3]int32{1,2,3}
	intArr := [...]int32{1, 2, 3} //shorthand implication for the variable of int32 of character length 3

	intSlice := []int32{4, 5, 6}
	fmt.Printf("The length is %v with the capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 8)
	fmt.Println(intSlice) //The sliced and appened value to the slice
	fmt.Printf("The length is %v with the capacity %v\n", len(intSlice), cap(intSlice))

	intSlice2 := []int32{10, 11, 12}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("The new sliced array with second slice as spread %v\n", intSlice)
	fmt.Printf("this is the memory location %v\n", &intArr[1]) // printing out the memory location of the entry
	fmt.Println(&intArr[2])
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("The length is %v with the capacity %v\n", len(intSlice3), cap(intSlice3))

	//MAP
	var myMap = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"John": 2, "Doe": 8}
	fmt.Println(myMap2["John"])

	var age, ok = myMap2["dason"]
	// delete(myMap2, "John")
	// fmt.Println(myMap2)

	if !ok {
		fmt.Println("The value is not found")
	} else {
		fmt.Println(age)
	}

	// first element is the key the second is value in a map
	for name, age := range myMap2 {
		fmt.Println(name, age)
	}

	// first elememnt is the index and the second is the value
	for i, v := range intSlice {
		fmt.Printf("The value %v is in the index %v\n", v, i)
	}

	//optional for while loop

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var n int = 1000000
	var slice1 []int = []int{}
	var slice2 []int = make([]int, 0, n)

	fmt.Println("The Slice without preallocation of memory", timeLoop(slice1, n))
	fmt.Println("The Slice with preallocation of memory", timeLoop(slice2, n))

	//String Manipulation
	var myStringRune = []rune("résume")
	var indexed = myStringRune[1]

	for i, v := range myStringRune {
		fmt.Println(i, v) // v is the value at that index, which will be a rune, or Unicode code point.
	}
	fmt.Println(indexed)

	var str = []string{"s", "u", "b", "s"}
	var strBuilder strings.Builder

	for i := range str {
		strBuilder.WriteString(str[i])
	}

	fmt.Printf("the concatenated string is %v\n", strBuilder.String())
	fmt.Printf("%b\n", myStringRune[1])

	// var gasEngineVar gasEngine = gasEngine{mpg: 10, gallons: 20, owner: owner{name: "Samar"}, brandName: brandName{"bmw"}, int: 10}

	// fmt.Println(gasEngineVar.mpg, gasEngineVar.gallons, gasEngineVar.owner.name, gasEngineVar.brandName.brandNameVar, gasEngineVar.int)
	var gasEngineVar gasEngine = gasEngine{10, 20, owner{name: "Samar"}, brandName{"bmw"}, 10}

	fmt.Println(gasEngineVar.mpg, gasEngineVar.gallons, gasEngineVar.owner.name, gasEngineVar.brandNameVar, gasEngineVar.int)
	fmt.Println("The miles left is", gasEngineVar.milesLeft())

	var electricEngineVar electricEngine = electricEngine{5, 15}

	// fmt.Println(canMakeit(gasEngineVar, 150))
	fmt.Println(canMakeit(electricEngineVar, 150))

	//POINTERS
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	p = &i //points the memory location of i to p
	*p = 1 //storing a value in that mem location, Now both p and i are pointing to same mem location so this var assignment will reflect to both *p and i

	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	var slice []int32 = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println("The result of both slice and sliceCopy as the underline behaviour is similar to pointer", slice, sliceCopy)

	var thing1 [5]float64 = [5]float64{1, 2, 3, 4, 5}
	fmt.Println("The memory location of the first char", &thing1[0])
	var results = square(&thing1)
	fmt.Println(thing1, results)
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner   owner
	brandName
	int
}

type owner struct {
	name string
}

type brandName struct {
	brandNameVar string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

// INTERFACE
type engine interface {
	milesLeft() uint8
}

func canMakeit(e engine, miles uint8) string {
	fmt.Println(e.milesLeft())
	if miles < e.milesLeft() {
		return "It can make it "
	} else {
		return "It can not make it "
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

func timeLoop(slice []int, n int) time.Duration {
	t0 := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	t1 := time.Now()
	return t1.Sub(t0)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Println("The memory location of the second char", &thing2[0])
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
