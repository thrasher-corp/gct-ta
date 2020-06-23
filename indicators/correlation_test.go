package indicators

import "testing"

func TestSum(t *testing.T) {
	if s := sum([]float64{1, 2, 3, 4}); s != 10 {
		t.Error("should return 10")
	}
}

func TestCorrelation(t *testing.T) {
	if r := correlation(nil, []float64{1234}); r != 0 {
		t.Error("should return 0")
	}

	n1 := []float64{41, 19, 23, 40, 55, 57, 33}
	n2 := []float64{94, 60, 74, 71, 82, 76, 61}
	if r := correlation(n1, n2); r != 0.5398442342154088 {
		t.Error("unexpected value")
	}
}

func TestCorrelationCoefficient(t *testing.T) {
	if r := CorrelationCoefficient(nil, nil, 5); r != nil {
		t.Error("invalid params should return nil")
	}
	closures1 := []float64{4086.29, 4310.01, 4509.08, 4130.37, 3699.99, 3660.02, 4378.48, 4640.00, 5709, 5950.02}
	closures2 := []float64{299.10, 348.13, 341.77, 293.50, 257.55, 282.00, 303.95, 311.07, 337.96, 293.79}
	if r := CorrelationCoefficient(closures1, closures2, 100); r != nil {
		t.Error("excess period should return nil")
	}
	if r := CorrelationCoefficient(closures1, closures2, -1); r != nil {
		t.Error("negative period should return nil")
	}
	if r := CorrelationCoefficient(closures1, closures2, 5); r[4] != 0.9370153710333355 || r[9] != 0.5294441817654052 {
		t.Error("unexpected result")
	}
	if r := CorrelationCoefficient(closures1, closures1, 5); r[4] != 1 {
		t.Error("same data should be 1")
	}
}

func BenchmarkCorrelationCoefficient(b *testing.B) {
	c1 := []float64{4086.29, 4310.01, 4509.08, 4130.37, 3699.99, 3660.02, 4378.48, 4640.00, 5709, 5950.02}
	c2 := []float64{299.10, 348.13, 341.77, 293.50, 257.55, 282.00, 303.95, 311.07, 337.96, 293.79}

	for x := 0; x < b.N; x++ {
		CorrelationCoefficient(c1, c2, 5)
	}
}
