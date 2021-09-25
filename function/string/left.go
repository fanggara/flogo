package string

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type Left struct {
}

func init() {
	function.Register(&Left{})
}

func (s *Left) Name() string {
	return "left"
}

func (s *Left) GetCategory() string {
	return "string"
}

func (s *Left) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt}, false
}

func (s *Left) Eval(params ...interface{}) (interface{}, error) {
	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Left first argument must be string")
	}
	size, err := coerce.ToInt(params[1])
	if err != nil {
		return nil, fmt.Errorf("Left second argument must be int")
	}

	if size < 0  {
		return nil, fmt.Errorf("Left second argument must be greater than 0")
	}

	if size == 0 {
		return "",nil
	}

	if len(str) < size {
		return str,nil 
	}

	return str[:len(str)-size-1],nil
	
}