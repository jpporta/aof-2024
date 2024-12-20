package p4

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
	if res != 18 {
		t.Fatalf("Answer not correct: %v", res)
	}
}

func TestFirstProblem(t *testing.T) {
	dat, err := os.ReadFile("./problem.txt")
	if err != nil {
		t.Fatal(err)
	}
	res := First(string(dat))
	if res != 2521 {
		t.Fatalf("Answer not correct %v", res)
	}
}

func TestSecondExample(t *testing.T) {
	dat, err := os.ReadFile("./example.txt")
	if err != nil {
		t.Fatal(err)
	}
	res := Second(string(dat))
	if res != 9 {
		t.Fatalf("Answer not correct: %v", res)
	}
}

func TestSecondProblem(t *testing.T) {
	dat, err := os.ReadFile("./problem.txt")
	if err != nil {
		t.Fatal(err)
	}
	res := Second(string(dat))
	if res != 1912 {
		t.Fatalf("Answer not correct: %v", res)
	}
}
