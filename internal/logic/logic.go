// Package logic implements core functionality of each REST endpoint
package logic

import (
	"log"

	"dispatch-auction/internal/share/model"
	"dispatch-auction/internal/share/util"
	"dispatch-auction/internal/storage"
)

type Logic struct {
	Database *storage.Database
}

// New is the constructor for the Logic struct
func New(database *storage.Database) *Logic {
	return &Logic{Database: database}
}

// CreateUser will create a user using the provided request data
func (l *Logic) CreateUser(request model.CreateUserRequest) (model.CreateUserResponse, error) {
	request.ID = util.NewID(16)
	err := l.Database.CreateUser(request)
	if err != nil {
		return model.CreateUserResponse{}, err
	}

	user := model.User{
		ID:       request.ID,
		Username: request.UserName,
	}
	return model.CreateUserResponse{User: user}, nil
}

// CreateAuction will create an auction using the provided request data
func (l *Logic) CreateAuction(request model.CreateAuctionRequest) (model.AuctionResponse, error) {
	request.ID = util.NewID(16)
	err := l.Database.CreateAuction(request)
	if err != nil {
		return model.AuctionResponse{}, err
	}

	auction := model.Auction{
		ID:          request.ID,
		OwnerID:     request.OwnerID,
		Title:       request.Title,
		Description: request.Description,
		Images:      request.Images,
	}
	return model.AuctionResponse{Auction: auction}, nil
}

// RegisterAuction will register a user as a bidder to an existing auction
func (l *Logic) RegisterAuction(request model.RegisterAuctionRequest) (model.RegisterAuctionResponse, error) {
	request.ID = util.NewID(16)
	err := l.Database.RegisterUserForAuction(request)
	if err != nil {
		return model.RegisterAuctionResponse{}, err
	}

	registration := model.AuctionRegistration{
		RegistrationID: request.ID,
		AuctionID:      request.AuctionID,
		BidderID:       request.BidderID,
	}
	return model.RegisterAuctionResponse{Registration: registration}, nil
}

// StartAuction will start an existing auction
func (l *Logic) StartAuction(request model.StartAuctionRequest) (model.AuctionResponse, error) {
	log.Printf("Auction %s is beginning\n", request.AuctionID)

	auction, err := l.Database.GetAuction(request.AuctionID, request.OwnerID)
	if err != nil {
		return model.AuctionResponse{}, err
	}

	bidders, err := l.Database.GetBiddersByAuctionID(auction.ID)
	if err != nil {
		return model.AuctionResponse{}, err
	}

	winner, winningBid := l.findWinner(auction, bidders)

	auction.WinningBid = winningBid
	auction.WinningUserID = &winner.User.ID
	return model.AuctionResponse{Auction: auction}, nil
}

// findWinner contains the main algorithm for the the Dispatch take-home challeng
// It will return the user with the winning bid given an auction and a list of bidding users
// Bids are served in round-robin fashion assuming the order in which bidders appear in the list will be the order in which bids are placed
func (l *Logic) findWinner(auction model.Auction, bidders []model.UserWithRegistrationData) (model.UserWithRegistrationData, int) {
	// prep bidders for auction by setting current bid to initial bid
	for i := range bidders {
		bidders[i].Registration.CurrentBid = bidders[i].Registration.InitialBid
	}

	idx := 0
	winningBid := 0
	for len(bidders) > 1 {
		bidder := bidders[idx]
		log.Printf("user %s bids %d\n", bidder.User.ID, bidder.Registration.CurrentBid)
		// check lose state: if current_bid > max_bid, remove bidder from slice
		if bidder.Registration.CurrentBid > bidder.Registration.MaxBid {
			log.Printf("user %s has exceeded the max bid of %d with %d\n", bidder.User.ID, bidder.Registration.CurrentBid, bidder.Registration.MaxBid)
			bidders = util.RemoveUserAtIndex(bidders, idx)
		} else {
			// check if this is the new winning bid
			if bidder.Registration.CurrentBid > winningBid {
				log.Printf("user %s is currently winning with a bid of %d\n", bidder.User.ID, bidder.Registration.CurrentBid)
				winningBid = bidder.Registration.CurrentBid
			}

			// auto-increment
			bidders[idx].Registration.CurrentBid += bidder.Registration.AutoIncrement
		}

		// increment index
		if idx+1 > len(bidders)-1 {
			idx = 0
		} else {
			idx += 1
		}
	}

	log.Printf("user %s wins at %d", bidders[0].User.ID, winningBid)
	return bidders[0], winningBid
}
