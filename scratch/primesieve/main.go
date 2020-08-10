package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	upperBound := "1234"
	//./primesieve 1 18446744073709551611 -n
	cmd := exec.Command("primesieve", "1",  upperBound, "-n")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf(err.Error())
	}
	output := out.String()
	if strings.Contains(output, "Nth prime:") {
		fmt.Print(output[strings.LastIndex(output, "Nth prime: ") + len("Nth prime: "):])
	}
}
