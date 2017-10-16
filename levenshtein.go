package main

// LevenshteinDistance measures the difference between two strings.
// The Levenshtein distance between two words is the minimum number of
// single-character edits (i.e. insertions, deletions or substitutions)
// required to change one word into the other.
//
// This implemention is optimized to use O(min(m,n)) space and is based on the
// optimized C version found here:
// http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C
func levenshteinDistance(r1, r2 string) int {
	//r1, r2 := []rune(s), []rune(t)
	var (
		oldDiag  int
		x        int
		y        int
		lastDiag int
		lenR1    = len(r1)
		lenR2    = len(r2)
	)

	column := make([]int, lenR1+1)

	for y = 1; y <= lenR1; y++ {
		column[y] = y
	}

	for x = 1; x <= lenR2; x++ {
		column[0] = x

		for y, lastDiag = 1, x-1; y <= lenR1; y++ {
			oldDiag = column[y]
			cost := 0
			if r1[y-1] != r2[x-1] {
				cost = 1
			}
			column[y] = min(column[y]+1, column[y-1]+1, lastDiag+cost)
			lastDiag = oldDiag
		}
	}

	return column[len(r1)]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}
