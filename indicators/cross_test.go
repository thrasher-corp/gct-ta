package indicators

import "testing"

var (
	testCrossunder1 = []float64{1, 2, 3, 4, 8, 6, 7}
	testCrossunder2 = []float64{1, 1, 10, 9, 5, 3, 7}

	testNothingCrossed1 = []float64{1, 2, 3, 4, 8, 6, 7}
	testNothingCrossed2 = []float64{1, 4, 5, 9, 5, 3, 7}
	testCrossover1      = []float64{1, 3, 2, 4, 8, 3, 8}
	testCrossover2      = []float64{1, 5, 1, 4, 5, 6, 7}
)

func TestCrossover(t *testing.T) {
	if Crossover(testCrossunder1, testCrossunder2) == true {
		t.Error("Crossover: Not expected and found")
	}

	if Crossover(testNothingCrossed1, testNothingCrossed2) == true {
		t.Error("Crossover: Not expected and found")
	}

	if Crossover(testCrossover1, testCrossover2) == false {
		t.Error("Crossover: Expected and not found")
	}
}
