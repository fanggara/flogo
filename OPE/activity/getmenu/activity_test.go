package getmenu

import (
	"net/http"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&OPEGetMenuActivity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &OPEGetMenuActivity{}
	act.json = jsoniter.ConfigCompatibleWithStandardLibrary
	act.setting = &Settings{OPEEndpoint: "http://192.168.23.20:8090/opes/api/v1/offers/eligible",TimeoutMs: 2000}
	act.client = &http.Client{Timeout: 2000 * time.Millisecond} 

	tc := test.NewActivityContext(act.Metadata())

	aInput := &Input{
		BusinessTransactionID: "tibcotest-trad-menu",
		RequestControl: map[string]interface{}{
			"returnProductInformationData" : map[string]interface{}{
				"returnCharacteristicsName" : "HEADER,DISPLAY_PRIORITY,COMMERCIAL_NAME",
				"returnCharacteristicsDescription" : "ProductComprisedOf",
				"returnCharacteristicsValueType" : "Instance,LANGUAGE,MENU",
			  	"returnCharacteristicsDetails" : false,
			},
			"validateData" : true,
			"validateProdDate" : true,
			"validateProductRequiredForGroups" : true,
			"validateProductComprisedOfGroups" : true,
			"validateProductCompatibility" : true,
			"validateSegmentCompatibility" : true,
			"enforceCompatibleSegment" : true,
			"skipChildFilterValidation" : true,
			"decomposeProducts" : true,
			"basicValidationOnExistingOffer" : true,
		},
		RecordTypes: map[string]interface{}{
			"recordTypes" : []string{"MENU"},
		},
		RecordSubTypes: map[string]interface{}{
			"recordSubTypes" : []string{"Business Product"},
		},
		Promotions: map[string]interface{}{
			"productIDs" : []string{"ML2_MENU_11"},
		},
		Segments: []interface{}{
			map[string]interface{}{
				"names" : []string{"Smart Phone" },
				"type" : "DEVICE_TYPE",
			},
		},
		ReturnIneligibleProducts: false,
		ReturnBundleOfferings: true,
		ReturnPrices: false,
		ReturnProductInformation: true,
	}

	tc.SetInputObject(aInput)
	done, err := act.Eval(tc)
	assert.NoError(t, err)
	assert.True(t,done)
	if err == nil{
		aOutput := &Output{}
		err = tc.GetOutputObject(aOutput)
		assert.NoError(t, err)
		assert.NotNil(t, aOutput.Menu)
		t.Log(aOutput)
	}
}