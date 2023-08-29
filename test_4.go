// test_4 - на коэффициент a равный 0 

go
func TestSolveAZero(t *testing.T) {
  _, err := solve(0, 1, 1)
  if err == nil {
    t.Errorf("expected error for a=0")
  }
}

