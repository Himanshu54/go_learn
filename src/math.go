package main

import (
	"fmt"
	"math/big"
  "math"
)

func main() {

	i1, i2, i3 := 12, 24, 48
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum: ", intSum)

	f1, f2, f3 := 12.2, 24.4, 48.8
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum: ", floatSum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(12.2)
	b2.SetFloat64(24.4)
	b3.SetFloat64(48.8)
	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("BigSum = %.10g\n", &bigSum)

  circleRadius := 15.5
  circumference := circleRadius * math.Pi
  fmt.Printf("Circumference : %.2f\n",circumference)
}
