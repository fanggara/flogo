package getmenu

import (
	"sort"
	"strconv"

	"github.com/beevik/etree"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&OPEMapMenuActivity{},New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}


const(
	elementNameUMB = "umb"
	elementNameTransactionID = "transaction_id"
	elementNameChannel =  "channel"
	elementNameStatusCode = "status_code"
	elementNameStatusDescription = "status_description"
	elementNamePath = "path"
	elementNameFirst = "first"
	elementNameMenu = "menu"
	elementNameNext = "next"
	elementNameBackCode = "backCode"
	elementAttrTotalMenu = "totalmenu"
	elementNameData = "data"
	elementAttrDataCode = "code"
	elementAttrDataPID = "pid"
	elementNameDataValue = "value"
	elementAttrDataPriority = "priority"
	dataCodeHeader = "0"
)

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &OPEMapMenuActivity{} //add aSetting to instance

	return act, nil
}

type menu struct {
	priority int
	value string
	pid string
	isHeader bool
}

// OPEGetMenuActivity is an sample OPEGetMenuActivity that can be used as a base to create a custom activity
type OPEMapMenuActivity struct {
}

// Metadata returns the activity's metadata
func (a *OPEMapMenuActivity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *OPEMapMenuActivity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	if ctx.Logger().DebugEnabled(){
		ctx.Logger().Infof("Input: %+v", input)
	}
	

	doc := etree.NewDocument()
	umbRoot := doc.CreateElement(elementNameUMB)
	umbRoot.CreateElement(elementNameTransactionID).SetText(input.TransactionID)
	umbRoot.CreateElement(elementNameChannel).SetText(input.Channel)
	umbRoot.CreateElement(elementNameStatusCode).SetText(input.StatusCode)
	umbRoot.CreateElement(elementNameStatusDescription).SetText(input.StatusDescription)
	umbRoot.CreateElement(elementNamePath).SetText(input.Path)
	umbRoot.CreateElement(elementNameFirst).SetText(input.First)

	if len(input.Menu) >0{
		menuElement := umbRoot.CreateElement(elementNameMenu)
		//totalMenu :=  0

		menuList := make([]menu, len(input.Menu))
		for i, m := range input.Menu {
			switch m.(type) {
			case map[string]interface{}:
				mCast := m.(map[string]interface{})
				menuItem := menu{}
				if priority,ok := mCast[elementAttrDataPriority];ok{
					prio,err := strconv.Atoi(priority.(string))
					if err != nil {
						return false, err
					}
					menuItem.priority = prio
				}
				if value,ok := mCast[elementNameDataValue];ok{
					menuItem.value  = value.(string)
				}
				if pid,ok := mCast[elementAttrDataPID];ok{
					menuItem.pid  = pid.(string)
				}
				menuList[i] = menuItem
			}
		}
		sort.Slice(menuList, func(i, j int) bool{
			return menuList[i].priority > menuList[j].priority
		})

		if input.Header != "" {
			data := menuElement.CreateElement(elementNameData)
			data.CreateAttr(elementAttrDataCode, dataCodeHeader)
			data.SetText(input.Header)
		}

		for i, m := range menuList {
			data := menuElement.CreateElement(elementNameData)
			data.CreateAttr(elementAttrDataCode, strconv.Itoa(i+1))
			data.CreateAttr(elementAttrDataPID, m.pid)
			data.SetText(m.value)
		}

		menuElement.CreateAttr(elementAttrTotalMenu, strconv.Itoa(len(input.Menu)))
	}

	umbRoot.CreateElement(elementNameNext).SetText(input.Next)
	umbRoot.CreateElement(elementNameBackCode).SetText(input.BackCode)
	
	result,err :=doc.WriteToString()
	if err != nil {
		return false, err
	}

	output := &Output{XMLContent: result}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	if ctx.Logger().DebugEnabled(){
		ctx.Logger().Debugf("Output: %+v", output)
	}

	return true, nil
}