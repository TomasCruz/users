package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertEqualError(t *testing.T, exp error, got error, args ...interface{}) {
	if exp == nil && got == nil {
		return
	}

	if exp == got {
		return
	}

	if exp == nil {
		assert.FailNow(t, "expected no error", "got: %s", got.Error())
		return
	}

	assert.EqualError(t, got, exp.Error(), args)
}
