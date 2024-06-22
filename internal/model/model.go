// Package model stores global models for the Auction application
package model

// Image is the model for optional images of auctioned items
type Image struct {
	URL  string `json:"url`
	Size struct {
		Width  int32
		Height int32
	}
	Resolution string
}
