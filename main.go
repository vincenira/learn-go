package main

import "fmt"

const (
	pi     = 3.14
	first  = iota
	second = iota + 6
	third  = 2 >> iota
)

const (
	fi = iota
	se
	th
)

func main() {
	var i int
	i = 42
	var fl float32 = 42.32
	mob_value := 3
	local_pi := 3.1453
	fmt.Println("Hello from a module, Gophers")
	fmt.Println(pi, first, second, third, fi, se, th)
	fmt.Println(i, fl, mob_value, local_pi, mob_value+3, float32(mob_value)+23.44)
}
