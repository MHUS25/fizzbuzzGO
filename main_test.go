package main

import "testing"

func TestFizz(t *testing.T) {
	if out := FizzBuzz(3); out != "Fizz" {
		t.Errorf("Expecting Fizz but got %v ", out)
	}
}

func TestFizz2(t *testing.T) {
	if out := FizzBuzz(9); out != "Fizz" {
		t.Errorf("Expecting Fizz but got %v ", out)
	}
}

func TestBuzz(t *testing.T) {
	if out := FizzBuzz(5); out != "Buzz" {
		t.Errorf("Expecting Buzz but got %v ", out)
	}
}

func TestBuzz2(t *testing.T) {
	if out := FizzBuzz(10); out != "Buzz" {
		t.Errorf("Expecting Buzz but got %v ", out)
	}
}

func TestFizzBuzz(t *testing.T) {
	if out := FizzBuzz(15); out != "FizzBuzz" {
		t.Errorf("Expecting FizzBuzz but got %v ", out)
	}
}

func TestFizzBuzz2(t *testing.T) {
	if out := FizzBuzz(30); out != "FizzBuzz" {
		t.Errorf("Expecting FizzBuzz but got %v ", out)
	}
}

func TestNumberOne(t *testing.T) {
	if out := FizzBuzz(1); out != "1" {
		t.Errorf("Expecting 1 but got %v ", out)
	}
}

func TestNumberSeven(t *testing.T) {
	if out := FizzBuzz(7); out != "7" {
		t.Errorf("Expecting 7 but got %v ", out)
	}
}
