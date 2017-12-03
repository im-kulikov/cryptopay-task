package internal

// Решение такое:
// - берём длину слова:
// - сначала проходим по схожей длине и вычисляем дистанцию по алгоритму Левенштейна
// - если результат нам не подходит - идём по шагам (слева / справа)
func (v Vocabulary) Distance(word string, count int, notify chan<- int) {
	// берём длину слова:
	wordLen := len(word)
	min := wordLen

	// сначала проходим по схожей длине
	// и вычисляем дистанцию по алгоритму Левенштейна
	if dist := v.partSearch(word, wordLen); dist == 1 {
		notify <- count
		return
	} else if min > dist {
		min = dist
	}

	// если результат нам не подходит - идём по шагам (слева / справа)
	step := 1
	for min > step {
		// Проверяем шаг вправо:
		if dist := v.stepSearch(word, wordLen+step, wordLen); dist == 1 {
			notify <- count
			return
		} else if min > dist {
			min = dist
		}

		// Проверяем шаг влево:
		if dist := v.stepSearch(word, wordLen-step, wordLen); dist == 1 {
			notify <- count
			return
		} else if min > dist {
			min = dist
		}

		step++
	}

	notify <- min * count
}
