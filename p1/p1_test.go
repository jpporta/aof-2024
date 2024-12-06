package p1

import (
	"os"
	"testing"
)

func TestSolveExample(t *testing.T) {
	dat, err := os.ReadFile("./example.txt")
	if err != nil {
		t.Fatal(err)
	}
	res := First(string(dat))
	if res != 11 {
		t.Fatalf("Answer not correct")
	}
}

func TestSolveProblem(t *testing.T) {
	dat, err := os.ReadFile("./problem.txt")
	if err != nil {
		t.Fatal(err)
	}
	res := First(string(dat))
	if res != 2285373 {
		t.Fatalf("Answer not correct")
	}
}

func TestSimilarityScore(t *testing.T) {
	dat, err := os.ReadFile("./example.txt")
	if err != nil {
		t.Fatal(err)
	}
	res := Second(string(dat))
	if res != 31 {
		t.Fatalf("Answer not correct: %v", res)
	}
}

func TestSimilarityProblem(t *testing.T) {
	dat, err := os.ReadFile("./problem.txt")
	if err != nil {
		t.Fatal(err)
	}
	res := Second(string(dat))
	if res != 21142653 {
		t.Fatalf("Answer not correct: %v", res)
	}
}
