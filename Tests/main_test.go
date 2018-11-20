package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "1.000000", fizzBuzz(1))
}
