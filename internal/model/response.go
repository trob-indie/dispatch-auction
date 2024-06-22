package model

import "time"

// CreateUserRequest is the model used to represent a response from to POST /api/user
type CreateUserResponse struct {
	ID        string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateUserResponse is the model used to represent a response from to POST /api/auction
type CreateAuctionResponse struct {
	ID          string
	OwnerID     string
	Title       string
	Description string
	Images      []Image
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// StartAuctionResponse is the model used to represent a response from to PUT /api/auction
type StartAuctionResponse struct {
	ID            string
	AuctionID     string
	OwnerID       string
	Title         string
	Description   string
	Images        []Image
	Bidders       []string
	WinningUserID string
	WinningBid    string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// RegisterAuctionResponse is the model used to represent a response from to POST /api/auction/register
type RegisterAuctionResponse struct {
	RegistrationID string
	AuctionID      string
	BidderID       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
