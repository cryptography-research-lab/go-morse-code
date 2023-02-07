package morse_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeToString(t *testing.T) {
	encode, err := EncodeToString("CC11001100")
	assert.Nil(t, err)
	assert.Equal(t, "-.-. -.-. .---- .---- ----- ----- .---- .---- ----- -----", encode)
}

func TestDecode(t *testing.T) {
	ciphertext := "-.-. -.-. .---- .---- ----- ----- .---- .---- ----- -----"
	decode, err := Decode(ciphertext)
	assert.Nil(t, err)
	assert.Equal(t, "CC11001100", decode)
}
