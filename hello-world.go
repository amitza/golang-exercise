package main

import (
	"fmt"
	"math"
	"time"
)

const s string = "constant"

func main() {
	name := "Go Developer"
	fmt.Println("Go for", name)

	//
	fmt.Println("go" + "lang")

	fmt.Println("1+1=", 1+1)

	fmt.Println("7.0/3.0=", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Vars
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	f := "apple"
	fmt.Println(f)

	// Consts
	fmt.Println(s)
	const n = 500000000
	const dd = 3e20 / n

	fmt.Println(dd)

	fmt.Println(int64(dd))

	fmt.Println(math.Sin(n))

	// For
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for nn := 0; nn <= 5; nn++ {
		if nn%2 == 0 {
			continue
		}
		fmt.Println(nn)
	}

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Print(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	ii := 2
	fmt.Print("Write ", ii, " as ")
	switch ii {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Printf("Dont know the type %T\n", t)

		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	var arr [5]int
	fmt.Println("emp:", arr)

	arr[4] = 100
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[4])

	fmt.Println("len:", len(arr))

	bArr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", bArr)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	slice := make([]string, 3)
	fmt.Println("emp:", slice)
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	fmt.Println("set: ", slice)
	fmt.Println("get: ", slice[2])
	fmt.Println("len: ", len(slice))
	slice = append(slice, "d")
	slice = append(slice, "e", "f")
	fmt.Println("apd:", slice)

	sliceCopy := make([]string, len(slice))
	copy(sliceCopy, slice)
	fmt.Println("cpy:", sliceCopy)
	

}
