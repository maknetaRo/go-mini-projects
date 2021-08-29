package calculator

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) float64 {
	return a / b
}

func DivideInt(a, b int) int {
	return a / b
}

func Reminder(a, b int) int {
	return a % b
}

func FeetToMeters(a float64) (float64) {
	meterSquare := a * 0.09290304
	return meterSquare
}
func MetersToFeet(a float64) (float64) {	
	feetArea := a / 0.09290304
	return  feetArea
}