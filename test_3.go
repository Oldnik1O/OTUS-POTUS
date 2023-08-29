//test_3 -  для уравнения x^2 + 2x + 1 = 0
go
func TestSolveDoubleRoot(t *testing.T) {
  roots, err := solve(1, 2, 1)
  if err != nil {
    t.Errorf("unexpected error: %v", err)
  }
  if len(roots) != 1 || roots[0] != -1 {
    t.Errorf("expected single root -1, got: %v", roots)
  }
}
