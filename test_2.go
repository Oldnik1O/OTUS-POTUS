// test_2 - для уравнения x^2 - 1 = 0
go
func TestSolveTwoRoots(t *testing.T) {
  roots, err := solve(1, 0, -1)
  if err != nil {
    t.Errorf("unexpected error: %v", err)
  }
  if len(roots) != 2 || (roots[0] != 1 && roots[1] != 1) || (roots[0] != -1 && roots[1] != -1) {
    t.Errorf("expected roots 1 and -1, got: %v", roots)
  }
}


