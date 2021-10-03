package getmenu

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/fanggara/flogo/OPE/activity/getmenu/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&OPEGetMenuActivity{},New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

const(
	headerContentType string = "application/json"
	mapFieldProducts string = "products"
	mapFieldID string = "id"
	mapFieldName string = "name"
	mapFieldType string = "type"
	mapFieldSubType string = "subType"
	mapFieldPCO string = "pco"
	mapFieldAttribute string = "attr"

	charTypeCharacteristic string = "Characteristic"
	charTypeProductComprisedOf string = "ProductComprisedOf"
)

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("OPE Endpoint: %s", s.OPEEndpoint)
	ctx.Logger().Debugf("OPE Timeout: %d", s.TimeoutMs)

	client := &http.Client{Timeout: time.Duration(s.TimeoutMs) * time.Millisecond}

	act := &OPEGetMenuActivity{client: client,setting: s,json: jsoniter.ConfigCompatibleWithStandardLibrary} //add aSetting to instance

	return act, nil
}

// OPEGetMenuActivity is an sample OPEGetMenuActivity that can be used as a base to create a custom activity
type OPEGetMenuActivity struct {
	client *http.Client
	setting *Settings
	json jsoniter.API
}

// Metadata returns the activity's metadata
func (a *OPEGetMenuActivity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *OPEGetMenuActivity) Eval(ctx activity.Context) (done bool, err error) {
	startTime := time.Now()
	defer func(){
		if done {
			ctx.Logger().Infof("Successfully Invoke OPE Get Menu. Completed in %d ms", time.Since(startTime).Milliseconds())
		}else{
			ctx.Logger().Errorf("Error occurred while invoking OPE Get Menu: %+v. Completed in %d ms", err,time.Since(startTime).Milliseconds())
		}	
	}()
	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return 
	}

	if ctx.Logger().DebugEnabled(){
		ctx.Logger().Debugf("Input: %+v", input)
	}

	data,err := a.json.Marshal(input)
	if err != nil {
		return 
	}

	res,err := a.GetMenu(data)
	if err != nil {
		return 
	}
	output := &Output{Menu: res}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return 
	}

	if ctx.Logger().DebugEnabled(){
		ctx.Logger().Debugf("Output: %+v", output)
	}
	done = true
	return 
}

func (a *OPEGetMenuActivity) GetMenu(reqRaw []byte)(interface{},error){
	bodyBuf := bytes.NewBuffer(reqRaw)
	resp,err := a.client.Post(a.setting.OPEEndpoint, headerContentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK{
		data,err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		getOfferResp := model.GetOfferResponse{}
		err = jsoniter.Unmarshal(data, &getOfferResp)
		if err != nil {
			return nil, err
		}

		menuList:= make(map[string]interface{},len(getOfferResp.EligibleProducts.EligibleProduct))
		for _,e := range getOfferResp.EligibleProducts.EligibleProduct {
			menu := make(map[string]interface{})
			menu[mapFieldID] = e.ProductInformation.Product.ProductID
			menu[mapFieldName] = e.ProductInformation.Product.ProductName
			menu[mapFieldType] = e.ProductInformation.Product.ProductType
			menu[mapFieldSubType] = e.ProductInformation.Product.ProductSubType
			pco := make([]interface{},0)
			attr:= make(map[string]string)
	
			for _,char := range e.ProductInformation.Characteristic {
				if char.Description  ==  charTypeProductComprisedOf{
					var pcoMap map[string]interface{}
					pcoMap = map[string]interface{}{ mapFieldID: char.Name}
					pco = append(pco, pcoMap)
				}else if char.Description  ==  charTypeCharacteristic {
					if len(char.Values.Value) >0 {
						attr[strings.ToLower(char.Name)] = char.Values.Value[0].Value
					}
				}
			}
			menu[mapFieldPCO] = pco
			menu[mapFieldAttribute] = attr
			menuList[e.ProductInformation.Product.ProductID] = menu
		}
	
		result := make([]interface{},0)
		for _,v := range menuList{
			pcoMap := v.(map[string]interface{})[mapFieldPCO]
			pcoItem := make([]interface{},0)
			for _,p := range pcoMap.([]interface{}){
				
				pidChild := p.(map[string]interface{})[mapFieldID].(string)
				if m,ok := menuList[pidChild];ok{
					pcoItem = append(pcoItem, m)
				}
			}
			if len(pcoItem)>0 {
				v.(map[string]interface{})[mapFieldPCO] = pcoItem
				result = append(result, v)
			}
		}

		return result,nil
	}
	
	return nil,errors.New(fmt.Sprintf("Failed to invoke OPE with HTTP status: %d", resp.StatusCode))
	
}