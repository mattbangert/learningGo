package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := &Saiyan{"Goku", 9000}
	goku.Super()
	fmt.Println(goku.Power)
}

func (s *Saiyan) Super() {
	s.Power += 10000
}
