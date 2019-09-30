package endpoint

import (
	"golang.org/x/net/context"

	dto "github.com/mdstella/go-users/endpoint/dto"
	service "github.com/mdstella/go-users/service"
	"github.com/go-kit/kit/endpoint"
)

//MakeUppercaseEndpoint - endpoint to uppercase the words sent by parameter.
func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return dto.UppercaseResponse{V: v, Err: err.Error()}, nil
		}
		return dto.UppercaseResponse{V: v, Err: ""}, nil
	}
}

//MakeCountEndpoint - endpoint to count the ammount of letters the words sent by paprameter have.
func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CountRequest)
		v := svc.Count(req.S)
		return dto.CountResponse{V: v}, nil
	}
}
