package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"math"
	"math/big"
	"os/exec"
	"strconv"
	"strings"
)
const (
	SegmentSize = 1024
)

// Bucket to use any other value.
type Bucket struct {
	// true value.
	value big.Int
	// is prime or not.
	isPrime bool
	// padding, to fit in the memory cache.
	// this structure already 40 bytes.
	// good to have some padding to fit in 64 bytes.
}

/*!
	This function to select the largest number with the upper bound value
 	is provided. Using sieve of eratosthenes to make sure we have the efficient
 	performant with O(nloglogn). And memory cost is O(n).

	instead of calculate from 2 -> n ( where is limited to 128 bit length for now).
	We can do better but need more computing memory, which is not a high demand for this phase.

	In order to save memory , We can break 2 - n into multiple smaller segments:
		+ first we trying select the largest prime number from segment Sn, Sn-1, Sn-2,..., S1.
		+ We can stop when touch the first element.

	Every segment containe a list of struct as below:
			segment struct {
				value big.Int
				prime bool
			}
	That struct required 40 bytes each. The cache memory usually design with 2^n Kb.

	algorithm Sieve of Eratosthenes is
    input: an integer n > 1.
    output: all prime numbers from 2 through n.

    let A be an array of Boolean values, indexed by integers 2 to n,
    initially all set to true.

    for i = 2, 3, 4, ..., not exceeding √n do
        if A[i] is true
            for j = i^2 + 0*i, i^2+ 1*i, i^2+2*i, i^2+3*i, ..., not exceeding n do
                A[j] := false

    return all i such that A[i] is true.

*/

func TakeLargestPrime128Bits( start *big.Int, end *big.Int ) (*big.Int, error) {
	if end.Cmp(start) != 1 {
		fmt.Printf("invalid start=%d, end=%d\n", start, end)
		return big.NewInt(0), errors.New("invalid input")
	}
	// no need to check if start < 2. since it should never happen.
	// We can the length of segment is always with 64 bits width.
	// fine to convert ton int64.
	isPrimes := make(map[string]bool, end.Sub(end, start).Int64())
	for prime := start; prime.Cmp(end) == -1 ; prime.Add(prime, big.NewInt(1)) {
		isPrimes[prime.String()] = true
	}
	// take the square of the upper bound value.
	squareValue := end.Abs(end.Sqrt(end))
	// looping the reverse order from the largest to 2.
	for i := squareValue.Sub(squareValue, big.NewInt(1)); i.Cmp(squareValue) == 1; i.Sub(i, big.NewInt(1)) {
		if isPrimes[i.String()] == true {
			// second loop is from j = 0 -> n
			j := big.NewInt(0)
			index := i.Mul(i,i).Add(i.Mul(i,i), i.Mul(i, j))
			for {
				if index.Cmp(end) !=  -1 {
					break
				}
				isPrimes[index.String()] = false
				j.Add(j, big.NewInt(1))
				index = i.Mul(i,i).Add(i.Mul(i,i), i.Mul(i, j))
			}
		}
	}
	for i := end.Sub(end, big.NewInt(1)); i.Cmp(start) == 1; i.Sub(i, big.NewInt(1)) {
		if isPrimes[i.String()] == true {
			return i, nil
		}
	}
	return big.NewInt(0), errors.New("no prime found")
}

// Take prime number less than or equal to 64 bit width.
func TakeLargestPrimeNumber64Bit( start uint64, end uint64 ) (uint64, error) {
	if end < start {
		fmt.Printf("invalid start=%d, end=%d\n", start, end)
		return 0, errors.New("invalid input")
	}
	// creating a bool map from start->end
	if start < 2 {
		// increment start up minimun value is 2.
		start = 2
	}
	isPrimes := make(map[uint64]bool, end - start)
	// initialized the rest of array to true.
	for prime := start; prime <= end; prime++ {
		isPrimes[prime] = true
	}
	// take the square of the upper bound value.
	squareValue := uint64(math.Abs(math.Sqrt(float64(end))))
	// looping the reverse order from the largest to 2.
	for i := squareValue - 1; i >= start; i-- {
		if isPrimes[i] == true {
			// second loop is from j = 0 -> n
			j := uint64(0)
			index := i*i + i*j
			for {
				if index >= end {
					break
				}
				isPrimes[index] = false
				j++
				index = i*i + i*j
			}
		}
	}
	for i := end - 1; i > start; i-- {
		if isPrimes[i] == true {
			return i, nil
		}
	}
	return 0, errors.New("no prime found")
}
func nativeSievePrime( low uint64, hight uint64) uint64  {
	cmd := exec.Command("primesieve ", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Errorf("Error %s", err.Error())
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
/*
void simpleSieve(int limit, vector<int>& prime)
{
    bool mark[limit + 1];
    memset(mark, false, sizeof(mark));

    for (int i = 2; i <= limit; ++i) {
        if (mark[i] == false) {
            // If not marked yet, then its a prime
            prime.push_back(i);
            for (int j = i; j <= limit; j += i)
                mark[j] = true;
        }
    }
}
*/

func SimpleSieve( limit uint64 ) uint64 {
	mark := make([]bool, limit + 1)
	for i := 0; i < len(mark); i++ {
		mark[i] = true
	}
	for i := uint64(2); i * i <= limit; i++ {
		if mark[i] == true {
			for j := i*i; j < limit; j += i {
				mark[j] = false
			}
		}
	}
	for i := limit; i > 2 ; i++ {
		if mark[i] == true{
			return i
		}
	}
	return uint64(0)
}

func TakeLargestPrimeNumber( upperBound string ) (string, error) {
	bigInteger := big.Int{}
	bigInteger.SetString(upperBound, 10)
	ctx, _ := context.WithCancel(context.Background())
	c,_ := cpu.InfoWithContext(ctx)
	fmt.Printf("Cache size is = %d\n",c[0].CacheSize)
	if bigInteger.BitLen() > 64 {
		fmt.Printf("Take less than 64 big length %d\n", bigInteger.BitLen())
		rawUpperBound := bigInteger.Uint64()
		for {
			var startSegment uint64
			if rawUpperBound < 2 {
				return strconv.FormatUint(0, 10), errors.New("Invalid input, expected greater than 2")
			} else if rawUpperBound == 2 {
				// if the input already 2 return 2.
				return strconv.Itoa(2), nil
			} else if rawUpperBound < SegmentSize {
				// if the upper bound is less than the segment size.
				// simple start from 0 -> segmentsize.
				startSegment = 0
			} else {
				startSegment = rawUpperBound - SegmentSize
			}
			prime, err := TakeLargestPrimeNumber64Bit( startSegment , rawUpperBound )
			if err == nil {
				fmt.Printf("the returned prime is %d\n", prime)
				return strconv.FormatUint(prime, 10), nil
			}
			// if there is no prime found, and the start already 0. then we return 2.
			if startSegment == 0 {
				return strconv.Itoa(2), nil
			// if there is some number left, continue to calculate.
			} else if startSegment < 1024 {
				rawUpperBound = startSegment
			} else {
				rawUpperBound = rawUpperBound - startSegment;
			}
		}
	} else {
		result := SimpleSieve(bigInteger.Uint64())
		return strconv.FormatUint(result, 10), nil
	}

	return strconv.Itoa(0), errors.New("unsupported input")
}


/*
func main() {
	a := int(123)
	b := uint64(123)
	c := "foo"
	f := big.Int{}

	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c))
	fmt.Printf("f: %T, %d\n", f, unsafe.Sizeof(f))
	ca := int64(18446744073709551614)
	cb := int64(ca - 25)
	if cb < 0 {
		fmt.Printf("Negative %d", cb)
	} else {
		fmt.Printf("Positive %d", cb)
	}
	reader := 	bufio.NewReader(os.Stdin)
	fmt.Println("Please input a number")
	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		text =  strings.Replace(text, "\n", "" , -1 )
		if text == "e" {
			break
		}
		start := time.Now()
		fmt.Printf("the value is %s\n", TakeLargestPrimeNumber64(text))
		fmt.Printf("Elapsed time is %s\n", time.Since(start))
	}

*/