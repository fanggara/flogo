package string_test

import (
	"testing"

	stringFunc "github.com/fanggara/flogo/function/string"
	"github.com/stretchr/testify/assert"
)

func TestSubstring(t *testing.T) {

	testCases := []struct {
		source    string
		start int
		end  int
		expected string
		isError bool
	}{
		{
			source:    "2017-04-10T22:17:32.000+0000",
			start:   0,
			end: 4,
			expected: "2017",
		},
		{
			source:    "Hello World", 
			start:   0,
			end: 5, 
			expected: "Hello",
		},
		{
			source:    "Hello World", 
			start:   6,
			end: 12, //end out of bound
			expected: "World",
		},
		{
			source:    "Hello World",
			start:   1,
			end: 1, //start =  end
			isError: true,
		},
		{
			source:    "Hello World",
			start:   2,
			end: 1, //start >  end
			isError: true,
		},
		{
			source:    "Hello World",
			start:   -1,
			end: 1, //start <0
			isError: true,
		},
		{
			source:    "Hello World",
			start:   0,
			end: -1, //start <0
			isError: true,
		},
	}

	n := &stringFunc.Substring{}

	for _, test := range testCases {
		v, err := n.Eval(test.source, test.start,test.end)
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