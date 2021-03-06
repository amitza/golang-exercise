package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"runtime"
	"sync"
	"testing"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func TestPlus(t *testing.T) {
	ans := plus(1, 1)
	if ans != 2 {
		t.Error("1 + 1 is 2!", ans)
	}
}

func waitGroup(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func workerPool(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func routine(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		fmt.Println("num of routines", runtime.NumGoroutine())
	}
}

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

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elemes []T
	for e := lst.head; e != nil; e = e.next {
		elemes = append(elemes, e.val)
	}
	return elemes
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't do 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with it"}
	}
	return arg + 3, nil
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

	var mmm = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(mmm))

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	routine("direct")
	go routine("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")

	messages := make(chan string, 2)
	go func() {
		messages <- "ping"
		messages <- "pong"
	}()
	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)

	done := make(chan bool, 1)
	go worker(done)
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case ms1 := <-c1:
			fmt.Println("received", ms1)
		case ms2 := <-c2:
			fmt.Println("recieved", ms2)
		}
	}

	c1 = make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 = make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

	messages = make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("recived message", msg)
	default:
		fmt.Println("No message received")
	}

	msg = "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	jobs := make(chan int, 5)
	done = make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recevied job", j)
			} else {
				fmt.Println("recived all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)

	const numJobs = 5
	jobs = make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go workerPool(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			waitGroup(i)
		}()
	}
	wg.Wait()

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("requests", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	dat, err := os.ReadFile("/home/aamita/git/golang-exercise/golang-exercise.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dat))

	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status", resp.Status)
	scanner := bufio.NewScanner(resp.Body)

	for i:= 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8090", nil)
}
