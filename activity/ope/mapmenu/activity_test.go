package getmenu

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&OPEMapMenuActivity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &OPEMapMenuActivity{}

	tc := test.NewActivityContext(act.Metadata())

	aInput := &Input{
		TransactionID: "tibcotest-trad-menu",
		Channel: "a1",
		StatusCode: "0000",
		StatusDescription: "Success",
		Path: "363|1",
		First: "yes",
		Header: "Hot Promo",
		Menu: []interface{}{
			map[string]interface{}{
				"priority" : "9",
				"value" : "Extra Kuota Harian",
				"pid" : "ML3_MENU_111",
			},
			map[string]interface{}{
				"priority" : "8",
				"value" : "Extra Kuota Berlangganan",
				"pid" : "ML3_MENU_112",
			},
		},
		Next: "MENU",
		BackCode: "9",
	}

	tc.SetInputObject(aInput)
	done, err := act.Eval(tc)
	assert.NoError(t, err)
	assert.True(t,done)
	if err == nil{
		aOutput := &Output{}
		err = tc.GetOutputObject(aOutput)
		assert.NoError(t, err)
		assert.NotNil(t, aOutput.XMLContent)
		t.Log(aOutput.XMLContent)
	}
}