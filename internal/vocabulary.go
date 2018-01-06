package internal

type Vocabulary map[int][]string

// Решение такое:
// - берём длину слова:
// - сначала проходим по схожей длине и вычисляем дистанцию по алгоритму Левенштейна
// - если результат нам не подходит - идём по шагам (слева / справа)
func (v Vocabulary) Distance(word string, count int, notify chan<- int) {
	var (
		wordLen = len(word)
		min     = wordLen
		step    int
	)

	for step = 0; min > step; step++ {
		// Проверяем шаг вправо:
		if dist := v.stepSearch(word, step, wordLen); dist == 1 {
			notify <- count
			return
		} else if min > dist {
			min = dist
		}
	}

	notify <- min * count
}

func (v Vocabulary) partSearch(words []string, word string, wordLen int) int {
	distMin := wordLen

	for i := 0; i < len(words); i++ {
		var s = words[i]
		if dist := levenshtein(word, s); dist == 1 {
			return 1
		} else if distMin > dist {
			distMin = dist
		}
	}

	return distMin
}

func (v Vocabulary) stepSearch(word string, step, wordLen int) int {
	var min = wordLen

	if words, ok := v[wordLen+step]; ok {
		if distMin := v.partSearch(words, word, wordLen); distMin < min {
			min = distMin
		}
	}

	if step > 0 {
		if words, ok := v[wordLen-step]; ok {
			if distMin := v.partSearch(words, word, wordLen); distMin < min {
				min = distMin
			}
		}
	}

	return min
}
