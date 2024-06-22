// Package logic implements core functionality of each REST endpoint
package logic

import (
	"dispatch-auction/internal/model"
)

// CreateUser will create a user using the provided request data
func CreateUser(request model.CreateUserRequest) (model.CreateUserResponse, error) {
	return model.CreateUserResponse{}, nil
}

// CreateAuction will create an auction using the provided request data
func CreateAuction(request model.CreateAuctionRequest) (model.CreateAuctionResponse, error) {
	return model.CreateAuctionResponse{}, nil
}

// StartAuction will start an existing auction
func StartAuction(request model.StartAuctionRequest) (model.StartAuctionResponse, error) {
	return model.StartAuctionResponse{}, nil
}

// RegisterAuction will register a user as a bidder to an existing auction
func RegisterAuction(request model.RegisterAuctionRequest) (model.RegisterAuctionResponse, error) {
	return model.RegisterAuctionResponse{}, nil
}
