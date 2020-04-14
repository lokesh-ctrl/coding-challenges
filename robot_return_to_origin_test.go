package challenges

import "testing"

func TestJudgeCircle(t *testing.T) {
	moves := "LRRLUD"
	ans := JudgeCircle(moves)

	if ans != true {
		t.Errorf("Sum was incorrect")
	}
}
