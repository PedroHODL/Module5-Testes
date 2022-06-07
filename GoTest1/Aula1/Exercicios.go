package main

import (
	"fmt"

	"github.com/PedroHODL/Module5-Testes.git/GoTest1/Aula1/calc"
	"github.com/PedroHODL/Module5-Testes.git/GoTest1/Aula1/divide"
	"github.com/PedroHODL/Module5-Testes.git/GoTest1/Aula1/ordering"
)

func main() {
	s := calc.Sub(7, 5)
	fmt.Println(s)

	list := []int{985467238, 902578946, 23856, 523765, 9604826, 1367541, 586318, 658724}
	fmt.Println(ordering.Sort(list))

	d, err := divide.Division(15, 5)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(d)

	df, err := divide.Division(52, 0)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(df)
}
