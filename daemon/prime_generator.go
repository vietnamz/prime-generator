package daemon

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/vietnamz/prime-generator/daemon/config"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type PrimeService struct {
	cfg *config.Config
}

// Constructor to initialize an prime generator srv.
func newPrimeService(config *config.Config) *PrimeService{
	return &PrimeService{
		config,
	}
}
// TakeLargestPrimes returns a largest prime number
// with an upper bound input provided.
// Input : upperBound 	: uint64
// Output: result		: uint64
// Limitation:
//				input < 2^64(18446744073709551615)
// Notes: if the result == 0, means we have an error.
func (p *PrimeService) TakeLargestPrimes(upperBound string) string {
	//./primesieve 1 18446744073709551611 -n
	result := strconv.Itoa(0)
	valid, _ := regexp.Match("^[0-9]*$", bytes.NewBufferString(upperBound).Bytes())
	if valid == false {
		logrus.Printf("no valid")
		return result
	}
	cmd := exec.Command("primesieve", "1",  upperBound, "-n")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		logrus.Errorf("Encounter an error %s", err)
		return result
	}
	output := out.String()
	if strings.Contains(output, "Nth prime:") == false {
		logrus.Errorf("No prime has been found : %s", output)
		return result
	}
	resultStr := output[strings.LastIndex(output, "Nth prime: ") + len("Nth prime: "):]
	resultStr = strings.TrimSpace(resultStr)
	valid, _ = regexp.Match("^[0-9]*$", bytes.NewBufferString(resultStr).Bytes())
	if valid == false {
		return result
	}
	return resultStr
}
