package p2

import (
	"os"
	"testing"
)

func TestFirstExample(t *testing.T) {
	dat, err := os.ReadFile("./example.txt")
	if err != nil {
		t.Fatal(err)
	}
	res := First(string(dat))
	if res != 2 {
		t.Fatalf("Answer not correct: %v", res)
	}
}

func TestFirstProblem(t *testing.T) {
	dat, err := os.ReadFile("./problem.txt")
	if err != nil {
		t.Fatal(err)
	}
	res := First(string(dat))
	if res != 534 {
		t.Fatalf("Answer not correct %v", res)
	}
}

func TestSecondExample(t *testing.T) {
	dat, err := os.ReadFile("./example.txt")
	if err != nil {
		t.Fatal(err)
	}
	res := Second(string(dat))
	if res != 4 {
		t.Fatalf("Answer not correct: %v", res)
	}
}

func TestSecondProblem(t *testing.T) {
	dat, err := os.ReadFile("./problem.txt")
	if err != nil {
		t.Fatal(err)
	}
	res := Second(string(dat))
	if res != 21142653 {
		t.Fatalf("Answer not correct: %v", res)
	}
}
