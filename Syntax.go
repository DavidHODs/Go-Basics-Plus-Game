package main

import (
	"fmt"
	"runtime"
)

// var (
// 	actorName string = "Elizabeth Sladen"
// 	companion string = "Sarah Jane Smith"
// 	doctorNumber int = 3
// 	season int = 11
// )

// var i int = 42

// func main() {
// 	// var j int
// 	// j = 100
// 	// s := 33
// 	// var m int = 22
// 	i := 42
// 	// fmt.Println(i, j, m)
// 	fmt.Printf("%v, %T\n", i, i)
// 	var j string = strconv.Itoa(i)
// 	fmt.Printf("%v, %T", j, j)
// }

// func main() {
// 	// var n bool = true
// 	n := 1 == 1
// 	m := 1 == 2
// 	fmt.Printf("%v, %T\n", n, n)
// 	fmt.Printf("%v, %T\n", m, m)
// }

// func main() {
// 	var n int16 = 32
// 	var m int16 = 32
// 	fmt.Println(n + m)
// }

// func main() {
// 	s := "This is a String"
// 	b := []byte(s)
// 	fmt.Printf("%v, %T\n", b, b)
// }

// func main() {
// 	var r rune = 'a'
// 	fmt.Printf("%v, %T\n", r, r)
// }

// func main() {
// 	// const myConst float64 = math.Sin(1.57)
// 	var myConst float64 = math.Sin(1.57)
// 	fmt.Printf("%v, %T\n", myConst, myConst)
// }

// func main() {
// 	// grades := [...]int{97, 66, 76, 55}
// 	var students [3]string
// 	fmt.Printf("Students: %v\n", students)
// 	students[0] = "Lisa"
// 	students[1] = "Ahmed"
// 	students[2] = "Arnold"
// 	fmt.Printf("Students: %v\n", students)
// 	fmt.Printf("Student#1: %v\n", students[1])
// 	fmt.Printf("The number of students: %v\n", len(students))
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	// b := &a
// 	// b[1] = 5
// 	fmt.Printf("Length: %v\n", len(a))
// 	fmt.Printf("Capacity: %v\n", cap(a))
// 	// fmt.Println(b)
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	b := a[:]
// 	c := a[:6]
// 	d := a[3:]
// 	e := a[3:6]
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)
// }

// func main() {
// 	a := make([]int, 3, 100)
// 	fmt.Println(a)
// 	fmt.Printf("Length: %v\n", len(a))
// 	fmt.Printf("Capacity: %v\n", cap(a))
// }

// func main() {
// 	a := []int{}
// 	fmt.Println(a)
// 	fmt.Printf("Length: %v\n", len(a))
// 	fmt.Printf("Capacity: %v\n", cap(a))
// 	a = append(a, 1)
// 	fmt.Println(a)
// 	fmt.Printf("Length: %v\n", len(a))
// 	fmt.Printf("Capacity: %v\n", cap(a))
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	// b := a[:len(a) - 1]
// 	b := append(a[:2], a[3:]...)
// 	fmt.Println((b))
// 	fmt.Println(a)
// }

// func main() {
// 	statePopulations := make(map[string]int)
// 	statePopulations = map[string]int {
// 		"California": 3546782892,
// 		"Texas": 546789262728,
// 		"Florida": 268900374,
// 		"New York": 769038448,
// 		"Ohio": 620979101,
// 	}
// 	// m := map[[3]int]string{}
// 	statePopulations["Georgia"] = 109283993
// 	delete(statePopulations, "Georgia")
// 	fmt.Println(statePopulations["Ohio"])
// 	fmt.Println(statePopulations["Georgia"])

// 	_, ok := statePopulations["Ohio"]
// 	fmt.Println(ok)
// 	fmt.Println(len(statePopulations))
// }

// type Doctor struct {
// 	Number int
// 	ActorName string
// 	Companions []string
// }

// func main() {
// 	aDoctor := Doctor{
// 		Number: 3,
// 		ActorName: "John Doe",
// 		Companions: []string {
// 			"Lee Grant",
// 			"Jo Grant",
// 		},
// 	}
// 	fmt.Println(aDoctor.Companions[1])
// }

// func main() {
// 	aDoctor := struct{name string}{
// 		name: "John Pierce",
// 	}
// 	anotherDoctor := aDoctor
// 	anotherDoctor.name = "Tom Baker"
// 	fmt.Println(aDoctor)
// 	fmt.Println(anotherDoctor)
// }

// type Animal struct {
// 	Name string
// 	Origin string
// }

// type Bird struct {
// 	Animal
// 	SpeedKPH float32
// 	CanFly bool
// }

// func main() {
// 	b := Bird{}
// 	b.Name = "Emu"
// 	b.Origin = "Australia"
// 	b.SpeedKPH = 43
// 	b.CanFly = false
// 	fmt.Println(b)
// 	fmt.Println(b.Name)
// }

// func main() {
// 	b := Bird {
// 		Animal: Animal{Name: "Emu", Origin: "Australia"},
// 		SpeedKPH: 87,
// 		CanFly: true,
// 	}
// 	fmt.Println(b)
// }

// type Animal struct {
// 	Name string `required max: "100"`
// 	Origin string
// }

// func main() {
// 	t := reflect.TypeOf(Animal{})
// 	field, _ := t.FieldByName("Name")
// 	fmt.Println(field.Tag)
// }

// func main() {
// 	if true {
// 		fmt.Println("The test is true")
// 	}
// }

// func main() {
// 	statePopulations := make(map[string]int)
// 	statePopulations = map[string]int {
// 		"California": 3546782892,
// 		"Texas": 546789262728,
// 		"Florida": 268900374,
// 		"New York": 769038448,
// 		"Ohio": 620979101,
// 	}
// 	if pop, ok := statePopulations["Florida"]; ok {
// 		fmt.Println(pop)
// 	}
// }

// func main() {
// 	number := 50
// 	guess := 30
// 	if guess < 1 || returnTrue() || guess > 100 {
// 		fmt.Println("The guess must be between 1 and 100")
// 	} else {
// 		if guess >=1 && guess <=100 {
// 		if guess < number {
// 		fmt.Println("Too Low")
// 	}
// 		if guess > number {
// 			fmt.Println("Too High")
// 		}
// 		if guess == number {
// 			fmt.Println("You got it..")
// 		}
// 	}
// 	}

// 	fmt.Println(number<=guess, number>=guess, number!=guess)
// }

// func returnTrue() bool {
// 	fmt.Println("Returning True")
// 	return true
// }

// func main() {
// 	switch i := 2+3; i {
// 	case 1, 5, 10:
// 		fmt.Println("One")
// 		// fallthrough
// 	case 2, 4, 6:
// 		fmt.Println("Two")
// 	case 3, 20:
// 		fmt.Println("Three")
// 		// fallthrough
// 	default:
// 		fmt.Println("Not One or Two")
// 	}
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// 	g := 0
// 	for ; g < 10; g++ {
// 		fmt.Println(g)
// 	}
// 	for j, m := 0, 0; j < 5; m, j = m+1, j+1 {
// 		fmt.Println(j, m)
// 	}
// 	b := 0
// 	for {
// 		fmt.Println(b)
// 		b++
// 		if b == 50 {
// 			break
// 		}
// 	}
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	for k, v := range s {
// 		fmt.Println(k,v)
// 	}
// 	for _, v := range s {
// 		fmt.Println(v)
// 	}
// }

// func main() {
// 	defer fmt.Println("Start")
// 	defer fmt.Println("Middle")
// 	defer fmt.Println("Finish")
// }

// func main() {
// 	res, err := http.Get("http://www.google.com/robots.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer res.Body.Close()
// 	robots, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s", robots)
// }

// func main() {
// 	a, b := 1, 0
// 	ans := a/b
// 	fmt.Println(ans)
// }

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello Go!"))
// 	})
// 	err := http.ListenAndServe(":8000", nil)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// func main() {
// 	var a int = 42
// 	var b *int = &a
// 	fmt.Println(a, *b)
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		sayMessage("Hello Go!!", i)
// 	}
// }

// func sayMessage(msg string, idx int) {
// 	fmt.Println(msg)
// 	fmt.Println("The value of the index is", idx)
// }

// func main() {
// 	s := sum(1, 2, 3, 4, 5)
// 	fmt.Println("The sum is ", s)
// }

// func sum(values ...int) (result int){
// 	fmt.Println(values)
// 	for _, v := range values {
// 		result += v
// 	}
// 	return
// }

// func divide(a, b float64) (float64, error){
// 	if b == 0.0 {
// 		return 0.0, fmt.Errorf("Cannot divide by zero")
// 	}
// 	return a/b, nil
// }

// func main() {
// 	d, err := divide(5.0, 0.0)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(d)
// }

// func main() {
// 	// for i := 0; i < 5; i++ {
// 	// 	func(i int) {
// 	// 		fmt.Println(i)
// 	// 	}(i)
// 	// }
// 	var f func() = func() {
// 		fmt.Println("Hello Go")
// 	}
// 	f()
// }

// func main() {
// 	g := greeter {
// 		greeting: "Hello",
// 		name: "Go",
// 	}
// 	g.greet()
// 	fmt.Println("The new name is:", g.name)
// }

// type greeter struct {
// 	greeting string
// 	name string
// }

// func (g *greeter) greet() {
// 	fmt.Println(g.greeting, g.name)
// 	g.name = ""
// }

// func main() {
// 	var w Writer = ConsoleWriter{}
// 	fmt.Println("Hello, Playground")
// }

// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type ConsoleWriter struct {}

// func (cw ConsoleWriter) Write(data []byte) (int, error) {
// 	n , err := fmt.Println(string(data))
// 	return n, err
// }

// func main() {
// 	go sayHello()
// 	time.Sleep(100 * time.Millisecond)
// }

// func sayHello() {
// 	fmt.Println("Hello")
// }

// var wg = sync.WaitGroup{}

// func main() {
// 	var msg = "Hello"
// 	wg.Add(1)
// 	go func(msg string) {
// 		fmt.Println(msg)
// 		wg.Done()
// 	}(msg)
// 	msg = "Goodbye"
// 	wg.Wait()
// }

func main() {
	runtime.GOMAXPROCS(100)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}