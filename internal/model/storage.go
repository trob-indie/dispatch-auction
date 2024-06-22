package model

import "time"

// User is the database model representing a user
type User struct {
	ID        string
	Username  string
	Password  string
	CreatedAt time.Time // TODO: may need a better time type for postgres
	UpdatedAt time.Time
}

// Auction is the database model representing an Auction
type Auction struct {
	ID            string
	OwnerID       string
	Title         string
	Description   string
	Images        []Image
	WinningUserID *string
	WinningBid    int32
	CreatedAt     time.Time // TODO: may need a better time type for postgres
	UpdatedAt     time.Time
}

// AuctionRegistration is the database model representing the relationship between an Auction, its owner, and bidders
type AuctionRegistration struct {
	RegistrationID string
	AuctionID      string
	BidderID       string
	CreatedAt      time.Time // TODO: may need a better time type for postgres
	UpdatedAt      time.Time
}

// Bid is the database model representing a user's bid on an auction
type Bid struct {
	ID            string
	AuctionID     string
	BidderID      string
	MaxBid        string
	CurrentBid    string
	AutoIncrement string
	CreatedAT     string
}
