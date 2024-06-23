// Package storage implements the postgres storage layer
package storage

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"dispatch-auction/internal/share/model"
)

// Storage defines the storage layer contract
type Storage interface {
	CreateUser(user model.CreateUserRequest) error
	CreateAuction(auction model.CreateAuctionRequest) error
	RegisterUserForAuction(auctionRegister model.RegisterAuctionRequest) error
	UpdateAuctionWithWinner(winning_user_id string, winning_bid int) (model.Auction, error)
	GetAuction(auctionID, ownerID string) (model.Auction, error)
	GetBiddersByAuctionID(auctionID string) ([]model.UserWithRegistrationData, error)
}

// Database implements the storage layer contract
type Database struct {
	*sql.DB
}

// New is the constructor for Database
func New(db *sql.DB) *Database {
	return &Database{DB: db}
}

// Migrate migrates the database for application usage
func (d *Database) Migrate(migrationFilePath string) error {
	file, err := os.ReadFile(migrationFilePath)
	if err != nil {
		return err
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		_, err = d.DB.Exec(request)
		if err != nil {
			return err
		}
	}

	return nil
}

const createUserQuery = `
	INSERT INTO users (id, username, password)
		VALUES ($1, $2, $3); 
`

// CreateUser will insert a user into the users table
func (d *Database) CreateUser(user model.CreateUserRequest) error {
	_, err := d.DB.Exec(createUserQuery, user.ID, user.UserName, user.Password)
	if err != nil {
		return err
	}

	log.Printf("user %s created successfully\n", user.ID)

	return nil
}

const createAuctionQuery = `
	INSERT INTO auctions (id, owner_id, title, description, images)
		VALUES ($1, $2, $3, $4, $5::jsonb); 
`

// CreateAuction will insert an auction into the auctions table
func (d *Database) CreateAuction(auction model.CreateAuctionRequest) error {
	_, err := d.DB.Exec(createAuctionQuery, auction.ID, auction.OwnerID, auction.Title, auction.Description, auction.Images)
	if err != nil {
		return err
	}

	log.Printf("auction %s created successfully by user %s\n", auction.ID, auction.OwnerID)

	return nil
}

const auctionRegisterQuery = `
	INSERT INTO auction_registrations (id, auction_id, bidder_id)
		VALUES ($1, $2, $3); 
`

// RegisterUserForAuction will insert an auction registration record into the auction_registrations table
func (d *Database) RegisterUserForAuction(auctionRegister model.RegisterAuctionRequest) error {
	_, err := d.DB.Exec(auctionRegisterQuery, auctionRegister.ID, auctionRegister.AuctionID, auctionRegister.BidderID)
	if err != nil {
		return err
	}

	log.Printf("user %s registered to auction %s successfully\n", auctionRegister.BidderID, auctionRegister.AuctionID)

	return nil
}

const getAuctionQuery = `
	SELECT * FROM auctions WHERE auction_id=$1 AND owner_id=$2 ORDER BY created_at DESC LIMIT 1;
`

func (d *Database) GetAuction(auctionID, ownerID string) (model.Auction, error) {
	result := d.DB.QueryRow(auctionRegisterQuery, auctionID, ownerID)

	auction := model.Auction{}
	err := result.Scan(&auction)
	if err != nil {
		return model.Auction{}, err
	}

	return model.Auction{}, nil
}

const getBiddersQuery = `
	SELECT * FROM auction_registrations WHERE auction_id=$1 
		LEFT JOIN users ON auction_registrations.bidder_id = users.id;
`

// GetBiddersByAuctionID will return bidders joined with auction registration data for the provided auctionID
func (d *Database) GetBiddersByAuctionID(auctionID string) ([]model.UserWithRegistrationData, error) {
	rows, err := d.DB.Query(auctionRegisterQuery, auctionID)
	if err != nil {
		return []model.UserWithRegistrationData{}, err
	}

	users := []model.UserWithRegistrationData{}
	for rows.Next() {
		var user model.UserWithRegistrationData
		err = rows.Scan(&user)
		if err != nil {
			return []model.UserWithRegistrationData{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
