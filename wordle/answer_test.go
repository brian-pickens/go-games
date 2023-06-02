package wordle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Guess_Match_ResultShouldBeTrue(t *testing.T) {
	answer := NewAnswer("abcde")
	result, _ := answer.Guess(([5]rune)([]rune("abcde")))
	assert.True(t, result)
}

func Test_Guess_NoMatch_ResultShouldBeFalse(t *testing.T) {
	answer := NewAnswer("bcdef")
	result, _ := answer.Guess(([5]rune)([]rune("abcde")))
	assert.False(t, result)
}

func Test_Guess_Match_GuessStateShouldBeAllCorrect(t *testing.T) {
	word := "abcde"
	answer := NewAnswer(word)
	expected := [5]state{ correct, correct, correct, correct, correct }
	_, guessState := answer.Guess(([5]rune)([]rune("abcde")))
	assert.ElementsMatch(t, expected, guessState)
}

func Test_Guess_PartialMatch_GuessStateShouldMatch(t *testing.T) {
	word := "abccc"
	answer := NewAnswer(word)
	expected := [5]state{ correct, correct, correct, incorrect, incorrect }
	_, guessState := answer.Guess(([5]rune)([]rune("abcde")))
	assert.ElementsMatch(t, expected, guessState)
}

func Test_Guess_WrongOrder_GuessStateShouldShowPresent(t *testing.T) {
	word := "aeghi"
	answer := NewAnswer(word)
	expected := [5]state{ correct, incorrect, incorrect, incorrect, present }
	_, guessState := answer.Guess(([5]rune)([]rune("abcde")))
	assert.ElementsMatch(t, expected, guessState)
}