package main

import "fmt"

func main() {
	a := [5]int{10, 20, 30, 40, 50}

	s := a[1:]

	fmt.Println("array:", a, "len:", len(a), "cap:", cap(a))
	fmt.Println("slices:", s, "len:", len(s), "cap:", cap(s))

	s[1] = 60

	fmt.Println("After change array:", a, "len:", len(a), "cap:", cap(a))
	fmt.Println("after change slices:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, 70)

	fmt.Println("After append array:", a, "len:", len(a), "cap:", cap(a))
	fmt.Println("after append slices:", s, "len:", len(s), "cap:", cap(s))

}
