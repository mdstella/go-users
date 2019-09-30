package encoder

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
)

//EncodeResponse - encode the string response.
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
