package main

import "fmt"

type s1 struct {
	x int
}

type s2 struct {
	x int
}

func main() {
	c := s1{x: 5}
	c1 := s2{x: 5}
	if c == c1 {
		fmt.Println("Yes")
	}
}
