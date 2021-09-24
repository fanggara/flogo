package getoffers

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
)

func TestMapStructure(t *testing.T){
	prods:= createProducts()
	result:= make(map[string]interface{}) 
	err :=mapstructure.Decode(prods,&result)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n",result)
}

func TestMapUnmarshal(t *testing.T){
	prods:= createProducts()
	result:= make(map[string]interface{}) 
	data,err := json.Marshal(prods)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n",result)
}

func BenchmarkMapStructure(b *testing.B){
	prods:= createProducts()
	for i := 0; i < b.N; i++ {
		result:= make(map[string]interface{}) 
		err :=mapstructure.Decode(prods,&result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshal(b *testing.B){
	prods:= createProducts()
	for i := 0; i < b.N; i++ {
		result:= make(map[string]interface{}) 
		
		data,err := json.Marshal(prods)
		if err != nil {
			b.Fatal(err)
		}

		err = json.Unmarshal(data, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}