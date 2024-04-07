package main

import "fmt"

/* deklarasi variable global */

var g = "value g\n"

var (
	h = "value h\n"
	i = "value i\n"
)

func main() {
	/* macam-macam deklarasi variable di go */

	a := "value a"
	fmt.Println(a)

	b, c := "value b,", 3
	fmt.Println(b, c)

	d := true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "value f"
	fmt.Println(f)

	// call variable global
	fmt.Println(g, h, i)
}
