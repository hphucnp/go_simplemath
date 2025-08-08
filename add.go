// Package gosimplemath provides simple mathematical operations.
//
// This package is designed for basic arithmetic functions.
package gosimplemath

import "golang.org/x/exp/constraints"

// Number is a type constraint that matches any integer or floating-point type.
type Number interface {
    constraints.Integer | constraints.Float
}

// Add returns the sum of two numbers.
//
// This function adds the two parameters together and returns the result.
// For more information about addition, see:
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
    return a + b
}