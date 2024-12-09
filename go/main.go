package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize/convex/lp"
)

func main() {
	c := []float64{1, 2, 0, 0}
	A := mat.NewDense(2, 4, []float64{1, 3, -1, 0, 2, 1, 0, -1})
	b := []float64{30, 40}

	obj, m, err := lp.Simplex(c, A, b, 0, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("x= %v y= %v obj= %v\n", m[0], m[1], obj)
}
