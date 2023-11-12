package main

import (
	"fmt"
)

type student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

func main() {
	st := new(student)
	fmt.Printf("%+v\n", st)

}
