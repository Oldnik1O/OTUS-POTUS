//test_5 -  для дискриминанта близкого к 0
go
roots, err := solve(1, 1.000000001, 0.000000001)

func TestSolveDoubleRoot(t *testing.T) {
  roots, err := solve(1, 2, 1)
  if err != nil {
    t.Errorf("unexpected error: %v", err)
  }
  if len(roots) != 1 || roots[0] != -1 {
    t.Errorf("expected single root -1, got: %v", roots)
  }
}

