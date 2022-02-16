package main

func distributeCandies(candyType []int) int {
	m := map[int]bool{}
	for _, c := range candyType {
		m[c] = true
	}
	if len(m) < len(candyType)/2 {
		return len(m)
	}
	return len(candyType) / 2
}
