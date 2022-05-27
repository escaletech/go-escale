package server

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Response http.ResponseWriter
	Request  *http.Request
}

// swagger:model
type ErrorResponse struct {
	Error string `json:"error"`
}

func New(Response http.ResponseWriter, r *http.Request) *Request {
	return &Request{Response, r}
}

//Decode request body
func (r *Request) BodyDecode(value interface{}) error {
	return json.NewDecoder(r.Request.Body).Decode(value)
}

//Response success
func (r *Request) WriteSuccess(body interface{}, status int) {
	if body == nil {
		r.Response.WriteHeader(status)
		return
	}

	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteHeader(status)
	json.NewEncoder(r.Response).Encode(body)
}

//Response error
func (r *Request) WriteError(body interface{}, err error, status int) {
	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteHeader(status)

	if body == nil {
		json.NewEncoder(r.Response).Encode(ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(r.Response).Encode(body)
}
