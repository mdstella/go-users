package decoder

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	dto "github.com/mdstella/go-users/endpoint/dto"
)

//DecodeUppercaseRequest - uppercase decoder.
func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request dto.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//DecodeCountRequest - count decoder
func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request dto.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
