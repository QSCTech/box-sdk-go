package box

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlePlainText(t *testing.T) {
	var wrongResult int
	err := handlePlainText([]byte(""), &wrongResult)
	assert.Error(t, err)
	assert.Equal(t, "param kind must be *bool", err.Error())

	var rightResult bool
	err = handlePlainText([]byte("true"), &rightResult)
	assert.Error(t, err)
	assert.Equal(t, "data content is unsupported: true", err.Error())
}
