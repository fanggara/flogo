package string

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type Substring struct {
}

func init() {
	function.Register(&Substring{})
}

func (s *Substring) Name() string {
	return "substring"
}

func (s *Substring) GetCategory() string {
	return "string"
}

func (s *Substring) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt,data.TypeInt}, false
}

func (s *Substring) Eval(params ...interface{}) (interface{}, error) {
	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Substring first argument must be string")
	}
	start, err := coerce.ToInt(params[1])
	if err != nil {
		return nil, fmt.Errorf("Substring second argument must be int")
	}
	end, err := coerce.ToInt(params[2])
	if err != nil {
		return nil, fmt.Errorf("Substring third argument must be int")
	}

	if start <0  || end<0 {
		return nil, fmt.Errorf("Substring start and end argument must be greater than 0")
	}

	if start >= end {
		return nil, fmt.Errorf("Substring start argument must be less than to end")
	}

	if len(str) < start {
		return nil, fmt.Errorf("Substring start argument must be less than length")
	}

	if len(str) < end {
		end = len(str)
	}

	return str[start:end], nil
}