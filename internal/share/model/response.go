package model

// CreateUserRequest is the model used to represent a response from to POST /api/user
type CreateUserResponse struct {
	User User `json:"user"`
}

// StartAuctionResponse is the model used to represent a response from to METHOD /api/auction
type AuctionResponse struct {
	Auction Auction `json:"auction"`
}

// RegisterAuctionResponse is the model used to represent a response from to POST /api/auction/register
type RegisterAuctionResponse struct {
	Registration AuctionRegistration `json:"registration"`
}
