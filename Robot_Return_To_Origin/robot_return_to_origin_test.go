package challenges

import "testing"

func TestJudgeCircle(t *testing.T) {
	moves := "LRRLUD"
	ans := JudgeCircle(moves)

	if ans != true {
		t.Errorf("Some weirdo!")
	}
}

func TestJudgeCircle2(t *testing.T) {
	moves := "LRRLUDL"
	ans := JudgeCircle(moves)

	if ans != false {
		t.Errorf("Some weirdo!")
	}
}
