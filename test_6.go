// test_6 - проверка коэффициентов на NaN

go
func TestSolveNaNCoefficients(t *testing.T) {
  _, err := solve(math.NaN(), 1, 1)
  if err == nil {
    t.Errorf("expected error for NaN coefficient")
  }
}

