package main

import (
	"fmt"
	repo "jedi6/pkg"
	"math"
)

// ESERCIZIO 4 DECLARATION
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I'm ", p.first, p.last, "and I am", p.age, "years old")
}

//ESERCIZIO 5 DECLARATION
type square struct {
	lato float64
}
type circle struct {
	raggio float64
}

func (s square) Area() float64 {
	return s.lato * s.lato
}
func (c circle) Area() float64 {
	return math.Pi * 2 * c.raggio
}

type shape interface {
	Area() float64
}

func info(s shape) {
	x := s.Area()
	fmt.Println(x)
}

//ESERCIZIO 7 DECLARATION
var pippo int
var f func()

func main() {

	//ESERCIZIO 1 (VEDI REPO PER DEF. FUNC)
	fmt.Printf("%v", repo.Foo())
	a, b := repo.Bar()
	fmt.Printf("\n%v\t%v\t%T", a, b, b)

	//ESERCIZIO 2
	xi := []int{10, 10, 10, 10}
	fmt.Printf("\n%v", repo.Somma(xi...))

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("\n%v\n", repo.Somma2(ii))

	//ESERCIZIO 3
	var c, d int = 10, 12
	defer fmt.Printf("\n%v", repo.DefSum(c, d))
	fmt.Println("dovrebbe stampare prima di somma")

	//ESERCIZIO 4

	p1 := person{
		first: "Yuri",
		last:  "Benelli",
		age:   38,
	}
	p1.speak()

	//ESERCIZIO 5
	cerchio := circle{
		raggio: 16.58,
	}
	quadrato := square{
		lato: 58.56,
	}
	info(cerchio)
	info(quadrato)

	//ESERCIZIO 6
	func() {
		fmt.Println("prova")
	}()

	//ESERCIO 7

	f := func() {
		fmt.Print("ciao")
	}
	f()
	g := func(s string) {
		//DOMANDA: come gestisco il return?

	}
	g("ciao")

	//ESERCIZIO 8

	t := repo.Ciao()()
	fmt.Println(t)
}
