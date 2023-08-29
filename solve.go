//solve
go
package main

import (
  "errors"
  "math"
  "testing"
)

// функция solve
func solve(a, b, c float64) ([]float64, error) {
  if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
    return nil, errors.New("coefficients cannot be NaN")
  }
  if math.Abs(a) < 1e-10 {
    return nil, errors.New("a cannot be zero")
  }
  discriminant := b*b - 4*a*c
  if discriminant < 0 {
    return nil, nil
  }
  if math.Abs(discriminant) < 1e-10 {
    return []float64{-b / (2 * a)}, nil
  }
  root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
  root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
  return []float64{root1, root2}, nil
}
