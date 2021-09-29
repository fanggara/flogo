package model

type GetOfferResponse struct {
	ResultStatus          ResultStatus     `json:"resultStatus"`
	EligibleProducts      EligibleProducts `json:"eligibleProducts"`
	BusinessTransactionID string           `json:"businessTransactionID"`
}
type ResultStatus struct {
	Deployment string `json:"deployment"`
	Service    string `json:"service"`
	Operation  string `json:"operation"`
	Component  string `json:"component"`
	Severity   string `json:"severity"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}
type Product struct {
	ProductID   string `json:"productID"`
	ProductName string `json:"productName"`
	ProductType string `json:"productType"`
	ProductSubType string `json:"productSubType"`
	Description string `json:"description"`
}
type Value struct {
	ValueType string `json:"valueType"`
	Value     string `json:"value"`
}
type Values struct {
	Value []Value `json:"value"`
}
type Characteristic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Values      Values `json:"values"`
}
type ProductInformation struct {
	Product        Product          `json:"product"`
	Characteristic []Characteristic `json:"characteristic"`
}
type EligibleProduct struct {
	ProductInformation ProductInformation `json:"productInformation"`
	//Udf                []interface{}      `json:"udf"`
	//Message            []interface{}      `json:"message"`
	//PriceItem          []interface{}      `json:"priceItem"`
}
type EligibleProducts struct {
	EligibleProduct []EligibleProduct `json:"eligibleProduct"`
}