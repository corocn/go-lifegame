package fizzbuzz

import "testing"

func TestFizzBuzz_1(t *testing.T) {
	actual := FizzBuzz(1)
	expected := "1"

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestFizzBuzz_2(t *testing.T) {
	actual := FizzBuzz(2)
	expected := "2"

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestFizzBuzz_3(t *testing.T) {
	actual := FizzBuzz(3)
	expected := "Fizz"

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestFizzBuzz_5(t *testing.T) {
	actual := FizzBuzz(5)
	expected := "Buzz"

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestFizzBuzz_15(t *testing.T) {
	actual := FizzBuzz(15)
	expected := "FizzBuzz"

	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}