package Controllers

import (
	"GoApp/Model"
	"GoApp/Services"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type SortRequest struct {
	Unsorted []int `json:"unsorted"`
}

type SortResponse struct {
	Unsorted []int `json:"unsorted"`
	Sorted []int `json:"sorted"`
}


func MakeSortEndpoint(sortSvc Services.Sort) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(SortRequest)
		sorted := sortSvc.Sort(req.Unsorted)
		return SortResponse{req.Unsorted, sorted}, nil
	}
}

func DecodeSortRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request SortRequest
	err := json.NewDecoder(r.Body).Decode(&request);
	if err != nil || request.Unsorted == nil {
		return nil, Model.BAD_REQUEST
	}
	return request, nil
}