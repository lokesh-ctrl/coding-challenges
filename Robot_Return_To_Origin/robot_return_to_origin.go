package challenges

func JudgeCircle(moves string) bool {
	position := make([]int, 2)
	for _, char := range moves {
		switch char {
		case 'L':
			position[0] = position[0] - 1
		case 'R':
			position[0] = position[0] + 1
		case 'U':
			position[1] = position[1] + 1
		case 'D':
			position[1] = position[1] - 1
		}
	}
	if position[0] == 0 && position[1] == 0 {
		return true
	}
	return false
}
