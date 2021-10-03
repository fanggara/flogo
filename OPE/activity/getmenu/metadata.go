package getmenu

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	OPEEndpoint	string `md:"opeEndpoint,required"`
	//TenantID string `md:"tenantID,required"`
	TimeoutMs 	int `md:"timeoutMs,required"`
}

type Input struct {
	BusinessTransactionID 		string 					`md:"businessTransactionID" json:"businessTransactionID"`
	RequestControl 				map[string]interface{} 	`md:"requestControl" json:"requestControl"`
	RecordTypes 				map[string]interface{} 	`md:"recordTypes" json:"recordTypes"`
	RecordSubTypes 				map[string]interface{} 	`md:"recordSubType" json:"recordSubType"`
	Promotions					map[string]interface{}	`md:"promotions" json:"promotions"`
	Segments 					interface{} 			`md:"segments" json:"segments"`
	ReturnIneligibleProducts 	bool  					`md:"returnIneligibleProducts" json:"returnIneligibleProducts"`
	ReturnBundleOfferings		bool  					`md:"returnBundleOfferings" json:"returnBundleOfferings"`
	ReturnPrices				bool  					`md:"returnPrices" json:"returnPrices"`
	ReturnProductInformation	bool  					`md:"returnProductInformation" json:"returnProductInformation"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	trxID, err := coerce.ToString(values["businessTransactionID"])
	if err != nil {
		return err
	}
	r.BusinessTransactionID = trxID

	requestControl,err := coerce.ToObject(values["requestControl"])
	if err != nil {
		return err
	}
	r.RequestControl = requestControl

	recordTypes,err := coerce.ToObject(values["recordTypes"])
	if err != nil {
		return err
	}
	r.RecordTypes = recordTypes

	recordSubTypes,err := coerce.ToObject(values["recordSubType"])
	if err != nil {
		return err
	}
	r.RecordSubTypes = recordSubTypes

	promotions,err := coerce.ToObject(values["promotions"])
	if err != nil {
		return err
	}
	r.Promotions = promotions

	segments,err := coerce.ToArray(values["segments"])
	if err != nil {
		return err
	}
	r.Segments = segments

	returnIneligibleProducts, err := coerce.ToBool(values["returnIneligibleProducts"])
	if err != nil {
		return err
	}
	r.ReturnIneligibleProducts = returnIneligibleProducts

	returnBundleOfferings, err := coerce.ToBool(values["returnBundleOfferings"])
	if err != nil {
		return err
	}
	r.ReturnBundleOfferings = returnBundleOfferings

	returnPrices, err := coerce.ToBool(values["returnPrices"])
	if err != nil {
		return err
	}
	r.ReturnPrices = returnPrices

	returnProductInformation, err := coerce.ToBool(values["returnProductInformation"])
	if err != nil {
		return err
	}
	r.ReturnProductInformation = returnProductInformation

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"businessTransactionID": r.BusinessTransactionID,
		"requestControl": r.RequestControl,
		"recordTypes": r.RecordTypes,
		"recordSubType": r.RecordSubTypes,
		"promotions": r.Promotions,
		"segments" : r.Segments,
		"returnIneligibleProducts": r.ReturnIneligibleProducts,
		"returnBundleOfferings": r.ReturnBundleOfferings,
		"returnPrices": r.ReturnPrices,
		"returnProductInformation": r.ReturnProductInformation,
	}
}

type Output struct {
	Menu interface{}  `md:"menu"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	menus, err := coerce.ToArray(values["menu"])
	if err != nil {
		return err
	}
	o.Menu = menus
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"menu": o.Menu,
	}
}