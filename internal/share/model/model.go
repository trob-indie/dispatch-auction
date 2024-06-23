// Package model stores global models for the Auction application
package model

// Image is the model for optional images of auctioned items
type Image struct {
	URL  string `json:"url"`
	Size struct {
		Width  int32 `json:"width"`
		Height int32 `json:"height"`
	}
	Resolution string `json:"resolution"`
}

// UserWithRegistrationData wraps the User and AuctionRegistration models
type UserWithRegistrationData struct {
	User         *User
	Registration *AuctionRegistration
}
