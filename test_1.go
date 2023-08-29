// test_1 - тест для x^2 + 1 = 0
func TestSolveNoRoots(t *testing.T) {
  roots, err := solve(1, 0, 1)
  if err != nil {
    t.Errorf("unexpected error: %v", err)
  }
  if len(roots) != 0 {
    t.Errorf("expected no roots, got: %v", roots)
  }
}
