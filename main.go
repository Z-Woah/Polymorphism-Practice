package main

import "fmt"

type Worker struct {
	hourlyWage int
	hours      int
}

type Robber struct {
	moneyStolen int
	timesStolen int
	caught      bool
}

func (w Worker) CalculateExpences() int {
	return w.hourlyWage * w.hours
}

func (r Robber) CalculateExpences() int {
	if r.caught {
		return 0
	}

	return r.moneyStolen * r.timesStolen
}

type Expence interface {
	CalculateExpences() int
}

func Calculate(e []Expence) int {
	total := 0
	for _, p := range e {
		total += p.CalculateExpences()
	}
	return total
}

func main() {

	w1 := Worker{
		100,
		10,
	}

	w2 := Worker{
		10,
		15,
	}

	r1 := Robber{
		100,
		3,
		false,
	}

	r2 := Robber{
		1000,
		3,
		true,
	}

	v := []Expence{w1, w2, r1, r2}
	total := Calculate(v)
	fmt.Printf("Your expences are $%d", total)
}
