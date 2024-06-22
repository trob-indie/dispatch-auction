// Package handler stores REST handler functionality
package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"dispatch-auction/internal/logic"
	"dispatch-auction/internal/model"
)

const apiRoot = "/api"

// SetupRESTHandlers initializes handler functions for each REST API route
func SetupRESTHandlers() {
	// POST
	http.HandleFunc(fmt.Sprintf("%s/user", apiRoot), handleUserRequest)
	// POST + PUT
	http.HandleFunc(fmt.Sprintf("%s/auction", apiRoot), handleAuctionRequest)
	// POST
	http.HandleFunc(fmt.Sprintf("%s/auction", apiRoot), handleAuctionRegisterRequest)
}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(contentType, ";")[0]))
		if mediaType != "application/json" {
			fmt.Fprintf(w, "{\"error\": \"Content-Type header is not application/json.\"}")
			return
		}
	}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	switch r.Method {
	case "POST":
		log.Println("handling POST /api/user")

		var userRequest model.CreateUserRequest
		err := dec.Decode(&userRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%w\"}", err)
			return
		}

		response, err := logic.CreateUser(userRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%w\"}", err)
			return
		}
		json.NewEncoder(w).Encode(response)
	default:
		fmt.Fprintf(w, "{\"error\": \"Only POST methods are supported.\"}")
	}
}

func handleAuctionRequest(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(contentType, ";")[0]))
		if mediaType != "application/json" {
			fmt.Fprintf(w, "{\"error\": \"Content-Type header is not application/json.\"}")
			return
		}
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	switch r.Method {
	case "POST":
		log.Println("handling POST /api/auction")
		var createAuctionRequest model.CreateAuctionRequest
		err := dec.Decode(&createAuctionRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%w\"}", err)
			return
		}

		response, err := logic.CreateAuction(createAuctionRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%w\"}", err)
			return
		}
		json.NewEncoder(w).Encode(response)
	case "PUT":
		log.Println("handling PUT /api/auction")

		var startAuctionRequest model.StartAuctionRequest
		err := dec.Decode(&startAuctionRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%w\"}", err)
			return
		}

		response, err := logic.StartAuction(startAuctionRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%w\"}", err)
			return
		}
		json.NewEncoder(w).Encode(response)
	default:
		fmt.Fprintf(w, "{\"error\": \"Only POST and PUT methods are supported.\"}")
	}
}

func handleAuctionRegisterRequest(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(contentType, ";")[0]))
		if mediaType != "application/json" {
			fmt.Fprintf(w, "{\"error\": \"Content-Type header is not application/json.\"}")
			return
		}
	}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	switch r.Method {
	case "POST":
		log.Println("handling POST /api/auction/register")

		var auctionRegisterRequest model.RegisterAuctionRequest
		err := dec.Decode(&auctionRegisterRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%w\"}", err)
			return
		}

		response, err := logic.RegisterAuction(auctionRegisterRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%w\"}", err)
			return
		}
		json.NewEncoder(w).Encode(response)
	default:
		fmt.Fprintf(w, "{\"error\": \"Only POST methods are supported.\"}")
	}
}
