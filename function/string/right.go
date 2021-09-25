package string

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type Right struct {
}

func init() {
	function.Register(&Right{})
}

func (s *Right) Name() string {
	return "right"
}

func (s *Right) GetCategory() string {
	return "string"
}

func (s *Right) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt}, false
}

func (s *Right) Eval(params ...interface{}) (interface{}, error) {
	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Right first argument must be string")
	}
	size, err := coerce.ToInt(params[1])
	if err != nil {
		return nil, fmt.Errorf("Right second argument must be int")
	}

	if size < 0  {
		return nil, fmt.Errorf("Right second argument must be greater than 0")
	}

	if len(str) < size {
		return str,nil 
	}

	return str[len(str)-size:],nil
	
}