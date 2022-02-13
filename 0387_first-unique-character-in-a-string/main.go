package main

func firstUniqChar(s string) int {
	arr := make([]int, 26)
	for _, c := range s {
		arr[c-'a'] += 1
	}
	for idx, c := range s {
		v := arr[c-'a']
		if v == 1 {
			return idx
		}
	}
	return -1
}
