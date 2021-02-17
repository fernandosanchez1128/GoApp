package Controllers

import (
	"GoApp/Services"
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
)

type UserResponse struct {
	Unsorted []int `json:"unsorted"`
	Sorted   []int `json:"sorted"`
}

func MakeUserEndpoint(userSvc Services.UserInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		return userSvc.GetUser(req)
	}
}

func DecodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	document := vars["id"]
	return document, nil
}
