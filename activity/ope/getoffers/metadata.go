package getoffers

type Settings struct {
	Uri		string		`md:"uri"`  	// The URI of the service to invoke
}

type Input struct {
	Content	interface{}	`md:"content"`	// The message content to send. This is only used in POST, PUT, and PATCH
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"content":     i.Content,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	i.Content = values["content"]

	return nil
}