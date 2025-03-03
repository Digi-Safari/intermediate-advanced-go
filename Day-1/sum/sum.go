package sum

// Add is an exported function which adds two numbers
// To export a func we make the first letter caps
func Add(a, b int) {
	doStuff(a, b)
}

var MultiplyResult int

func Multiply(a, b int) {
	MultiplyResult = a * b
}
