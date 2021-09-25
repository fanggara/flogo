package string

import (
	"fmt"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

type Pad struct {
}

func init() {
	function.Register(&Pad{})
}

func (s *Pad) Name() string {
	return "Pad"
}

func (s *Pad) GetCategory() string {
	return "string"
}

func (s *Pad) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt}, false
}

func (s *Pad) Eval(params ...interface{}) (interface{}, error) {
	str, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Pad first argument must be string")
	}
	length, err := coerce.ToInt(params[1])
	if err != nil {
		return nil, fmt.Errorf("Pad second argument must be int")
	}
	justify, err := coerce.ToString(params[2])
	if err != nil {
		return nil, fmt.Errorf("Pad third argument must be string")
	}

	var format string 
	if strings.ToLower(justify) == "left" {
		format = fmt.Sprintf("%%-%vv", length)
	}else{
		format = fmt.Sprintf("%%0%vv", length)
	}
	fmt.Println(format)

	return fmt.Sprintf(format, str),nil
	
}