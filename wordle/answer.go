package wordle

type answer [COLUMNS]rune
type state int

const (
	empty state = iota
	incorrect
	present
	correct
)

func NewAnswer(str string) (answer) {
	an := ([5]rune)([]rune(str))
	return answer{an[0], an[1], an[2], an[3], an[4]}
}

func (an *answer) Guess(guess [COLUMNS]rune) (result bool, guessState [COLUMNS]state) {
	result = false
	if guess == ([5]rune)(*an) {
		result = true
	}

	for a := 0; a < COLUMNS; a++ {
		if result {
			guessState[a] = correct
			continue
		}

		if guess[a] == rune(0) {
			guessState[a] = empty
			continue
		}

		if guess[a] == an[a] {
			guessState[a] = correct
			continue
		}

		for b := 0; b < COLUMNS; b++ {
			if guess[a] == an[b] {
				guessState[a] = present
				return
			}
		}

		guessState[a] = incorrect
	}

	return result, guessState
}
