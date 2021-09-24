package getoffers

import (
	"encoding/json"

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
	data,err := json.Marshal(prods)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(data, &results)
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

type Eligible struct{
	Products []Product `json:"products"`
}

type Product struct{
	ID string `json:"id"`
	Type string `json:"type"`
	PCO []Product `json:"pco"`
}

func createProducts() (prods Eligible){
	prods = Eligible{Products: make([]Product, 0)}

	prod := Product{ID: "002992",Type: "PO",PCO: []Product{{ID: "MR_123",Type: "MR"}}}
	prods.Products = append(prods.Products, prod)
	return
}