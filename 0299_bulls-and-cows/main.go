package main

import "fmt"

func getHint(secret string, guess string) string {
	req := make([]int, 10)
	avail := make([]int, 10)
	bulls := 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			req[secret[i]-'0']++
			avail[guess[i]-'0']++
		}
	}
	cows := 0
	for i := 0; i < 10; i++ {
		if req[i] > 0 && avail[i] > 0 {
			cows += min(req[i], avail[i])
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
