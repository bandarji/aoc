package main

import (
	"fmt"
)

type Data struct {
	a, b, c 	int
}


func func1(d *Data) {
	d.a = 1
}

func func2() Data {
	d := Data{5, 5, 5}
	return d
}

func func3(d *Data) {
	d.a = 100
}

func main() {
	d := Data{0, 0, 0}
	func1(&d)
	fmt.Println(d)
	e := func2()
	fmt.Println(e)
	func3(&e)
	fmt.Println(e)
}
