package string_test

import (
	"testing"

	stringFunc "github.com/fanggara/flogo/function/string"
	"github.com/stretchr/testify/assert"
)

func TestRight(t *testing.T) {

	testCases := []struct {
		source    string
		size int
		expected string
		isError bool
	}{
		{
			source:    "Hello World", 
			size:   0, //return nothing
			expected: "",
		},
		{
			source:    "Hello World", 
			size:   5, //return last 5
			expected: "World",
		},
		{
			source:    "Hello World", 
			size:   12, //return all, size out of bound
			expected: "Hello World",
		},
		{
			source:    "Hello World", 
			size:   -2, //return all, size out of bound
			isError: true,
		},
		
	}

	n := &stringFunc.Right{}

	for _, test := range testCases {
		v, err := n.Eval(test.source, test.size)
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