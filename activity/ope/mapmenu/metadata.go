package getmenu

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
}

type Input struct {
	TransactionID 		string 			`md:"transactionID"`
	Channel 			string 			`md:"channel"`
	StatusCode 			string 			`md:"statusCode"`
	StatusDescription 	string 			`md:"statusDescription"`
	Path				string			`md:"path"`
	First				string			`md:"first"`
	Header				string			`md:"header"`
	Menu 				[]interface{} 	`md:"menu"`
	Next 				string  		`md:"next"`
	BackCode			string  		`md:"backCode"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	trxID, err := coerce.ToString(values["transactionID"])
	if err != nil {
		return err
	}
	r.TransactionID = trxID

	channel,err := coerce.ToString(values["channel"])
	if err != nil {
		return err
	}
	r.Channel = channel

	statusCode,err := coerce.ToString(values["statusCode"])
	if err != nil {
		return err
	}
	r.StatusCode = statusCode

	statusDescription,err := coerce.ToString(values["statusDescription"])
	if err != nil {
		return err
	}
	r.StatusDescription = statusDescription

	path,err := coerce.ToString(values["path"])
	if err != nil {
		return err
	}
	r.Path = path

	first,err := coerce.ToString(values["first"])
	if err != nil {
		return err
	}
	r.First = first

	header,err := coerce.ToString(values["header"])
	if err != nil {
		return err
	}
	r.Header = header

	menu,err := coerce.ToArray(values["menu"])
	if err != nil {
		return err
	}
	r.Menu = menu

	next, err := coerce.ToString(values["next"])
	if err != nil {
		return err
	}
	r.Next = next

	backCode, err := coerce.ToString(values["backCode"])
	if err != nil {
		return err
	}
	r.BackCode = backCode

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"transactionID": r.TransactionID,
		"channel": r.Channel,
		"statusCode": r.StatusCode,
		"statusDescription": r.StatusDescription,
		"path": r.Path,
		"first" : r.First,
		"header" : r.Header,
		"menu": r.Menu,
		"next": r.Next,
		"backCode": r.BackCode,
	}
}

type Output struct {
	XMLContent string  `md:"xmlContent"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	xmlContent, err := coerce.ToString(values["xmlContent"])
	if err != nil {
		return err
	}
	o.XMLContent = xmlContent
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"xmlContent": o.XMLContent,
	}
}