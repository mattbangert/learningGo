package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := new(Saiyan)
	goku.Name = "matt"
	goku.Power = 101
	goku.Super()
	fmt.Println(goku.Name, " ", goku.Power)
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}
