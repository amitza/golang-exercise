package main

import (
	"fmt"
	"math"
	"time"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func totalSum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func fib(n int) int {
	if n < 2 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 27
	return &p
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

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
	l := slice[2:5]
	fmt.Println("sl1:", l)
	l = slice[:5]
	fmt.Println("sl2:", l)
	l = slice[2:]
	fmt.Println("sl3:", l)

	tt := []string{"g", "h", "i"}
	fmt.Println("dcl:", tt)
	twoDim := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDim[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDim[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoDim)

	mm := make(map[string]int)
	mm["k1"] = 7
	mm["k2"] = 13
	fmt.Println("map: ", mm)
	v1 := mm["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len: ", len(mm))
	delete(mm, "k2")
	fmt.Println("map: ", mm)
	_, prs := mm["k2"]
	fmt.Println("prs:", prs)
	nn := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", nn)

	nums := []int{1, 2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println("index: ", i)
	}
	fmt.Println("sum: ", sum)
	kvs := map[string]string{"a": "tent", "b": "house"}
	for k, v := range kvs {
		fmt.Println("key:", k, "value: ", v)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	aa, bb := vals()
	fmt.Println(aa, bb)
	_, cc := vals()
	fmt.Println(cc)
	fmt.Println("sum is", totalSum(1, 2, 3, 4, 5))

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println(fib(7))

	i = 1
	zeroval(i) // by value
	fmt.Println(i)
	zeroptr(&i) // by reference
	fmt.Println(i)
	fmt.Println(&i)

	fmt.Println(person{"bob", 20})
	fmt.Println(newPerson("bob"))

	rec := rect{width: 10, height: 6}
	fmt.Println(rec.area(), rec.perim())
	recPointer := &rec
	fmt.Println(recPointer.area(), recPointer.perim())

	measure(rec)

	co := container{
		base: base{num: 1},
		str:  "some name",
	}

	type describer interface {
		describe() string
	}

	var ddd describer = co
	fmt.Println("desc: ", ddd.describe())
}
