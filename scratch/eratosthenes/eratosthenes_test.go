package eratosthenes

import (
	"testing"
)

/*
func TestTakeLargestPrimeNumber(t *testing.T) {

	got, err := TakeLargestPrimeNumber("12")
	if err != nil {
		t.Fatalf("return error with %s; want 11", got)
	} else {
		if got != "11" {
			t.Errorf("got %s; want 11", got)
		}
	}

}

func TestTakeLargestPrimeNumberWith2(t *testing.T) {
	got, err := TakeLargestPrimeNumber("2")
	if err != nil {
		t.Fatalf("return error with %s; want 2", got)
	} else {
		if got != "2" {
			t.Errorf("got %s; want 2", got)
		}
	}
}

func TestTakeLargestPrimeNumberWithless1024(t *testing.T) {
	got, err := TakeLargestPrimeNumber("56")
	if err != nil {
		t.Fatalf("return error with %s; want 53", got)
	} else {
		if got != "53" {
			t.Errorf("got %s; want 53", got)
		}
	}
}

func TestTakeLargestPrimeNumberlargestThan1024(t *testing.T) {
	got, err := TakeLargestPrimeNumber("1021")
	if err != nil {
		t.Fatalf("return error with %s; want 1019", got)
	} else {
		if got != "1019" {
			t.Errorf("got %s; want 1019", got)
		}
	}
}

func TestTakeLargestPrimeNumberlargestThan1024v2(t *testing.T) {
	got, err := TakeLargestPrimeNumber("7920")
	if err != nil {
		t.Fatalf("return error with %s; want 7919", got)
	} else {
		if got != "7919" {
			t.Errorf("got %s; want 7919", got)
		}
	}
}

func TestTakeLargestPrimeNumber1(t *testing.T) {
	_, err := TakeLargestPrimeNumber("1")
	if err == nil {
		t.Fatal("should be an error")
	}
}

func TestTakeLargestPrimeNumberVeryLargest(t *testing.T) {
	got, err := TakeLargestPrimeNumber("123456789012")
	if err != nil {
		t.Fatalf("return error with %s; want 123456789011;", got)
	} else {
		if got != "123456789011" {
			t.Errorf("got %s; want 123456789011;", got)
		}
	}
}

func TestTakeLargestPrimeNumberVeryLargest2(t *testing.T) {
	got, err := TakeLargestPrimeNumber("123456802730")
	if err != nil {
		t.Fatalf("return error with %s; want 123456802830", got)
	} else {
		if got != "123456802837" {
			t.Errorf("got %s; want 123456802837", got)
		}
	}
}
 */

func TestSimpleSieve(t *testing.T) {
	got := SimpleSieve(12)
	t.Fatalf("The return value is %d", got)
	t.Error("no value")
}