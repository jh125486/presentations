package main

import (
	"fmt"
)

// START OMIT
type (
	Celsius    float64 // Celsius represent temperature in celsius // HL
	Fahrenheit float64 // Fahrenheit represent temperature in fahrenheit
)

func (c Celsius) convert() Fahrenheit { // HL
	return Fahrenheit(c*(9/5.0) + 32) // HL
} // HL

func (c Celsius) Freezing() bool { // HL
    return c <= 0
} // HL

func main() {
	temperature := Celsius(20)
	fmt.Println(temperature.convert(), "F")
    fmt.Println("Freezing?:", temperature.Freezing())

    temperature = Celsius(-10)
    fmt.Println("Freezing?:", temperature.Freezing())
}
// END OMIT
