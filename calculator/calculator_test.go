package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2,2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel() 
	var want float64 = 6
	got := calculator.Multiply(3, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel() 
	var want float64 = 2
	got := calculator.Divide(6, 3)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivideInt(t *testing.T) {
	division := calculator.DivideInt(8, 2)
	expected := 4

	if division != expected {
		t.Errorf("expected '%d' but got '%d", expected, division)
	}
}

func TestReminder(t *testing.T) {
	reminder := calculator.Reminder(32, 9)
	expected := 5 

	if reminder != expected {
		t.Errorf("expected '%d' but got '%d", expected, reminder)
	}
}

func TestFeetToMeters(t *testing.T) {
	t.Parallel()
	var want float64 = 0.18580608
	got := calculator.FeetToMeters(2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMetersToFeet(t *testing.T) {
	t.Parallel()
	var want float64 = 107.639104
	got := calculator.MetersToFeet(10)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}