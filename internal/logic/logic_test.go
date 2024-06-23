package logic

import (
	"testing"

	"dispatch-auction/internal/share/model"
)

func Test_findWinner(t *testing.T) {

	type expected struct {
		winner     model.UserWithRegistrationData
		winningBid int
	}

	type provided struct {
		auction model.Auction
		bidders []model.UserWithRegistrationData
	}

	type test struct {
		provided provided
		expected expected
	}

	tests := map[string]test{
		"auction #1": {
			provided: provided{
				auction: model.Auction{
					ID:      "A1A1A1A1A1A1A1A1",
					OwnerID: "0000000000000000",
					Title:   "Auction #1",
				},
				bidders: []model.UserWithRegistrationData{
					{
						User: &model.User{
							ID:       "U1U1U1U1U1U1U1U1",
							Username: "Sasha",
						},
						Registration: &model.AuctionRegistration{
							RegistrationID: "R1R1R1R1R1R1R1R1",
							InitialBid:     50_00,
							MaxBid:         80_00,
							AutoIncrement:  3_00,
						},
					},
					{
						User: &model.User{
							ID:       "U2U2U2U2U2U2U2U2",
							Username: "John",
						},
						Registration: &model.AuctionRegistration{
							RegistrationID: "R2R2R2R2R2R2R2R2",
							InitialBid:     60_00,
							MaxBid:         82_00,
							AutoIncrement:  2_00,
						},
					},
					{
						User: &model.User{
							ID:       "U3U3U3U3U3U3U3U3",
							Username: "Pat",
						},
						Registration: &model.AuctionRegistration{
							RegistrationID: "R3R3R3R3R3R3R3R3",
							InitialBid:     55_00,
							MaxBid:         85_00,
							AutoIncrement:  5_00,
						},
					},
				},
			},
			expected: expected{
				winner: model.UserWithRegistrationData{
					User: &model.User{
						ID:       "U2U2U2U2U2U2U2U2",
						Username: "John",
					},
					Registration: &model.AuctionRegistration{
						RegistrationID: "R2R2R2R2R2R2R2R2",
						InitialBid:     60_00,
						MaxBid:         82_00,
						AutoIncrement:  2_00,
					},
				},
				winningBid: 85_00,
			},
		},
		"auction #2": {
			provided: provided{
				auction: model.Auction{
					ID:      "A1A1A1A1A1A1A1A1",
					OwnerID: "0000000000000000",
					Title:   "Auction #1",
				},
				bidders: []model.UserWithRegistrationData{
					{
						User: &model.User{
							ID:       "U1U1U1U1U1U1U1U1",
							Username: "Riley",
						},
						Registration: &model.AuctionRegistration{
							RegistrationID: "R1R1R1R1R1R1R1R1",
							InitialBid:     700_00,
							MaxBid:         725_00,
							AutoIncrement:  2_00,
						},
					},
					{
						User: &model.User{
							ID:       "U2U2U2U2U2U2U2U2",
							Username: "Morgan",
						},
						Registration: &model.AuctionRegistration{
							RegistrationID: "R2R2R2R2R2R2R2R2",
							InitialBid:     599_00,
							MaxBid:         725_00,
							AutoIncrement:  15_00,
						},
					},
					{
						User: &model.User{
							ID:       "U3U3U3U3U3U3U3U3",
							Username: "Charlie",
						},
						Registration: &model.AuctionRegistration{
							RegistrationID: "R3R3R3R3R3R3R3R3",
							InitialBid:     625_00,
							MaxBid:         725_00,
							AutoIncrement:  8_00,
						},
					},
				},
			},
			expected: expected{
				winner: model.UserWithRegistrationData{
					User: &model.User{
						ID:       "U3U3U3U3U3U3U3U3",
						Username: "Charlie",
					},
					Registration: &model.AuctionRegistration{
						RegistrationID: "R3R3R3R3R3R3R3R3",
						InitialBid:     625_00,
						MaxBid:         725_00,
						AutoIncrement:  8_00,
					},
				},
				winningBid: 724_00,
			},
		},
		"auction #3": {
			provided: provided{
				auction: model.Auction{
					ID:      "A1A1A1A1A1A1A1A1",
					OwnerID: "0000000000000000",
					Title:   "Auction #1",
				},
				bidders: []model.UserWithRegistrationData{
					{
						User: &model.User{
							ID:       "U1U1U1U1U1U1U1U1",
							Username: "Alex",
						},
						Registration: &model.AuctionRegistration{
							RegistrationID: "R1R1R1R1R1R1R1R1",
							InitialBid:     2500_00,
							MaxBid:         3000_00,
							AutoIncrement:  500_00,
						},
					},
					{
						User: &model.User{
							ID:       "U2U2U2U2U2U2U2U2",
							Username: "Jesse",
						},
						Registration: &model.AuctionRegistration{
							RegistrationID: "R2R2R2R2R2R2R2R2",
							InitialBid:     2800_00,
							MaxBid:         3100_00,
							AutoIncrement:  201_00,
						},
					},
					{
						User: &model.User{
							ID:       "U3U3U3U3U3U3U3U3",
							Username: "Drew",
						},
						Registration: &model.AuctionRegistration{
							RegistrationID: "R3R3R3R3R3R3R3R3",
							InitialBid:     2501_00,
							MaxBid:         3200_00,
							AutoIncrement:  247_00,
						},
					},
				},
			},
			expected: expected{
				winner: model.UserWithRegistrationData{
					User: &model.User{
						ID:       "U3U3U3U3U3U3U3U3",
						Username: "Drew",
					},
					Registration: &model.AuctionRegistration{
						RegistrationID: "R3R3R3R3R3R3R3R3",
						InitialBid:     2501_00,
						MaxBid:         3200_00,
						AutoIncrement:  247_00,
					},
				},
				winningBid: 3001_00,
			},
		},
	}

	l := Logic{}
	for name, test := range tests {
		t.Run(name, func(*testing.T) {
			actualWinner, actualWinningBid := l.findWinner(test.provided.auction, test.provided.bidders)

			if test.expected.winner.User.ID != actualWinner.User.ID {
				t.Errorf("expected winner to be %s, but got %s", test.expected.winner.User.ID, actualWinner.User.ID)
			}

			if test.expected.winningBid != actualWinningBid {
				t.Errorf("expected winning bid to be %d, but got %d", test.expected.winningBid, actualWinningBid)
			}
		})
	}
}
