// addint package provides ability to add two integers.
//
// This is a simple package that provides a function to add two integers.
// https://www.mathsisfun.com/numbers/addition.html
package addint

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// AddInt returns the sum of two integers.
//
// Example:
//
//	fmt.Println(AddInt(1, 2)) // Output: 3
func AddInt[T Number](a, b T) T {
	return a + b
}
