package getoffers

import (
	"github.com/mitchellh/mapstructure"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&MyActivity{},New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: %s", s.ASetting)

	act := &MyActivity{} //add aSetting to instance

	return act, nil
}

// MyActivity is an sample MyActivity that can be used as a base to create a custom activity
type MyActivity struct {
}

// Metadata returns the activity's metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *MyActivity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	ctx.Logger().Infof("TransactionID: %s", input.BusinessTransactionID)
	ctx.Logger().Infof("Segments: %v", input.Segments)

	prods := createProducts()
	results := make(map[string]interface{})
	err = mapstructure.Decode(prods, &results)
	if err != nil {
		return false, err
	}
	output := &Output{Products: results}

	ctx.Logger().Infof("Output: %v", output)
	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}

type eligible struct{
	products []product 
}

type product struct{
	id string 
	recordType string 
	pco []product
}

func createProducts() (prods eligible){
	prods = eligible{products: make([]product, 0)}

	prod := product{id: "002992",recordType: "PO",pco: []product{{id: "MR_123",recordType: "MR"}}}
	prods.products = append(prods.products, prod)
	return
}