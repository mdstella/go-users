package dto

//UppercaseRequest - request to uppercase words.
type UppercaseRequest struct {
	S string `json:"s"`
}

//UppercaseResponse - response to uppercase words.
type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

//CountRequest - request to count letters in words.
type CountRequest struct {
	S string `json:"s"`
}

//CountResponse - response to count letters in words.
type CountResponse struct {
	V int `json:"v"`
}
