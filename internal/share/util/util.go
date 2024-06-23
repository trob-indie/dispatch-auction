// Package util stores globally shared utiltity functions
package util

import (
	"time"

	"golang.org/x/exp/rand"

	"dispatch-auction/internal/share/model"
)

func Init() {
	rand.Seed(uint64(time.Now().UnixNano()))
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// NewID will return a random string of provided length
func NewID(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RemoveUserAtIndex(slice []model.UserWithRegistrationData, s int) []model.UserWithRegistrationData {
	return append(slice[:s], slice[s+1:]...)
}
