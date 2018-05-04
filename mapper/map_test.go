package mapper

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	mp := NewMap()
	mp.Put("1", "3")
	assert.Equal(t, 1, mp.Size())
	mp.Put("1", "3")
	assert.Equal(t, 1, mp.Size())
	mp.Put("2", "3")
	assert.Equal(t, 2, mp.Size())
}

func testHashFunction(key string) int {
	return 1
}
func TestOptions(t *testing.T) {
	mp := NewMap(WithHashFunction(testHashFunction), WithCapacity(10))
	assert.NotPanics(t, func() { mp.Put("2", "3") })
}
func TestMap(t *testing.T) {
	mp := NewMap()
	mp.Put("1", "3")
	mp.Put("2", "4")
	mp.Put("23", "411")
	mp.Put("123", "411")
	cases := []struct {
		key           string
		expectedValue string
		expectedError error
	}{
		{
			key:           "1",
			expectedValue: "3",
		},
		{
			key:           "2",
			expectedValue: "4",
		},
		{
			key:           "23",
			expectedValue: "411",
		},
		{
			key:           "3",
			expectedError: errors.New("key not found"),
		},
	}
	for _, c := range cases {
		actualValue, err := mp.Get(c.key)
		if c.expectedError != nil {
			assert.NotNil(t, err)
			assert.EqualError(t, err, c.expectedError.Error())
		} else {
			assert.Equal(t, c.expectedValue, actualValue)
			assert.Nil(t, err)
		}
	}
}
