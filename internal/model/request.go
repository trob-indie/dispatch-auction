package model

// CreateUserRequest is the model generated from a request to POST /api/user
type CreateUserRequest struct {
	ID       string
	UserName string
	Password string
}

// CreateAuctionRequest is the model generated from a request to POST /api/auction
type CreateAuctionRequest struct {
	OwnerID     string
	Title       string
	Description string
	Images      []Image
}

// StartAuctionRequest is the model generated from a request to PUT /api/auction
type StartAuctionRequest struct {
	AuctionID string
	OwnerID   string
}

// RegisterAuctionRequest is the model generated from a request to PUT /api/auction
type RegisterAuctionRequest struct {
	AuctionID string
	BidderID  string
}
