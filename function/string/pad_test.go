package string_test

import (
	"testing"

	stringFunc "github.com/fanggara/flogo/function/string"
	"github.com/stretchr/testify/assert"
)

func TestPadZero(t *testing.T) {

	testCases := []struct {
		source    string
		length int
		justify string
		expected string
		isError bool
	}{
		{
			source:    "12345", 
			length:   0, //return same
			expected: "12345",
		},
		{
			source:    "12345", 
			length:   10, //pad left 
			justify: "right",
			expected: "0000012345",
		},
		{
			source:    "12345", 
			length:   6, //pad left 
			justify: "left",
			expected: "12345 ",
		},
	}

	n := &stringFunc.Pad{}

	for _, test := range testCases {
		v, err := n.Eval(test.source, test.length,test.justify)
		if test.isError{
			assert.NotNil(t,err)
			assert.Nil(t, v)
			t.Log(err)
		}else{
			assert.Nil(t, err)
			assert.NotNil(t, v)
			assert.Equal(t, test.expected, v)
		}
		
	}
}