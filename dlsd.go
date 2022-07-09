package dlsd

func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func isSame(lhs, rhs []rune) bool {
	for i := 0; i < len(lhs); i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}

func Distance(lhs, rhs []rune) int {
	rl1, rl2 := len(lhs), len(rhs)

	if rl1 == 0 || rl2 == 0 {
		return max(rl1, rl2)
	}
	if rl1 == rl2 && isSame(lhs, rhs) {
		return 0
	}

	da := make(map[rune]int)
	for i := 0; i < rl1; i++ {
		da[lhs[i]] = 0
	}
	for i := 0; i < rl2; i++ {
		da[rhs[i]] = 0
	}

	d := make([][]int, rl1+2, rl1+2)
	for i := 0; i <= rl1+1; i++ {
		d[i] = make([]int, rl2+2, rl2+2)
		for j := 0; j <= rl2+1; j++ {
			d[i][j] = 0
		}
	}

	maxdist := rl1 + rl2
	d[0][0] = maxdist
	for i := 0; i <= rl1; i++ {
		d[i+1][0] = maxdist
		d[i+1][1] = i
	}
	for i := 0; i <= rl2; i++ {
		d[0][i+1] = maxdist
		d[1][i+1] = i
	}

	var cost int
	for i := 1; i <= rl1; i++ {
		db := 0
		for j := 1; j <= rl2; j++ {
			i1 := da[rhs[j-1]]
			j1 := db
			if lhs[i-1] == rhs[j-1] {
				cost = 0
				db = j
			} else {
				cost = 1
			}
			adddel := min(d[i+1][j]+1, d[i][j+1]+1)
			submov := min(d[i][j]+cost, d[i1][j1]+(i-i1-1)+1+(j-j1-1))
			d[i+1][j+1] = min(adddel, submov)
		}
		da[lhs[i-1]] = i
	}

	return d[rl1+1][rl2+1]
}
