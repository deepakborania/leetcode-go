package main

func maximumScore(a int, b int, c int) int {
	var x, y, z, score int

	// Set x as max number
	if a >= b && a >= c {
		x = a
		y = b
		z = c
	} else if b >= a && b >= c {
		x = b
		y = a
		z = c
	} else if c >= a && c >= b {
		x = c
		y = a
		z = b
	} else {
		x = a
		y = b
		z = c
	}

	for y+z >= x+2 {
		y--
		z--
		score++
	}
	for x > 0 && y > 0 {
		x--
		y--
		score++
	}
	for x > 0 && z > 0 {
		x--
		z--
		score++
	}
	return score
}
