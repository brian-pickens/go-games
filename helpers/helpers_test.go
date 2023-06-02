package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Min(t *testing.T) {
	a := 1
	b := 2
	assert.Equal(t, Min(a, b), a)
}

func Test_Max(t *testing.T) {
	a := 1
	b := 2
	assert.Equal(t, Max(a, b), b)
}