package model

import "time"

// User is the database model representing a user
type User struct {
	ID        string    `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"-"` // TODO: may need a better time type for postgres
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

// Auction is the database model representing an Auction
type Auction struct {
	ID            string    `db:"id" json:"id"`
	OwnerID       string    `db:"owner_id" json:"ownerID"`
	Title         string    `db:"title" json:"title"`
	Description   string    `db:"description" json:"description"`
	Images        []Image   `db:"images" json:"images"`
	WinningUserID *string   `db:"winning_user_id" json:"winningUserID"`
	WinningBid    int       `db:"winning_bid" json:"winningBid"`
	Bidders       []string  `db:"-" json:"bidders"`
	CreatedAt     time.Time `db:"created_at" json:"-"` // TODO: may need a better time type for postgres
	UpdatedAt     time.Time `db:"updated_at" json:"-"`
}

// AuctionRegistration is the database model representing the relationship between an Auction, its owner, and bidders
type AuctionRegistration struct {
	RegistrationID string    `db:"registration_id" json:"registrationID"`
	AuctionID      string    `db:"auction_id" json:"auctionID"`
	BidderID       string    `db:"bidder_id" json:"bidderID"`
	MaxBid         int       `db:"max_bid" json:"maxBid"`
	InitialBid     int       `db:"initial_bid" json:"initialBid"`
	CurrentBid     int       `db:"current_bid" json"current_bid"`
	AutoIncrement  int       `db:"auto_increment" json:"autoIncrement"`
	CreatedAt      time.Time `db:"created_at" json:"-"` // TODO: may need a better time type for postgres
	UpdatedAt      time.Time `db:"updated_at" json:"-"`
}
