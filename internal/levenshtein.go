package internal

// LevenshteinDistance measures the difference between two strings.
// The Levenshtein distance between two words is the minimum number of
// single-character edits (i.e. insertions, deletions or substitutions)
// required to change one word into the other.
//
// This implementation is optimized to use O(min(m,n)) space and is based on the
// optimized C version found here:
// http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C
var (
	levenshtein = levenshteinDistance
)

//func SetLevenshtein(method func(r1, r2 string) int) { levenshtein = method }

func levenshteinDistance(r1, r2 string) int {
	var (
		oldDiag  int
		x        int
		y        int
		p        int
		lastDiag int
		lenR1    = len(r1)
		lenR2    = len(r2)
	)

	for ; p < lenR1 && p < lenR2; p++ {
		if r2[p] != r1[p] {
			break
		}
	}
	r1, r2 = r1[p:], r2[p:]
	lenR1 -= p
	lenR2 -= p

	for 0 < lenR1 && 0 < lenR2 {
		if r1[lenR1-1] != r2[lenR2-1] {
			r1, r2 = r1[:lenR1], r2[:lenR2]
			break
		}
		lenR1--
		lenR2--
	}

	if lenR1 == 0 {
		return lenR2
	}

	if lenR2 == 0 {
		return lenR1
	}

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

	return column[lenR1]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}
