package main

import (
	"fmt"
	"math/big"
)
var wheel = [] *big.Int{
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),

}
wheel = {

}
func main() {
	a := big.NewInt(0)
	a.SetString("98357987892374589276895729067498690358609358093890683096803", 10)
	fmt.Println(a.ProbablyPrime(256))
}
