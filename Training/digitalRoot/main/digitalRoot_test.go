package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	assert.Equal(t, 7, digitalRoot(16))
	assert.Equal(t, 6, digitalRoot(942))
	assert.Equal(t, 6, digitalRoot(132189))
	assert.Equal(t, 2, digitalRoot(493193))

}
