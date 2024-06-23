// Package handler stores REST handler functionality
package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"dispatch-auction/internal/logic"
	"dispatch-auction/internal/share/model"
)

const apiRoot = "/api"

type handler struct {
	logic *logic.Logic
}

// SetupRESTHandlers initializes handler functions for each REST API route
func SetupRESTHandlers(logic *logic.Logic) {
	h := handler{logic: logic}

	// POST
	http.HandleFunc(fmt.Sprintf("%s/user", apiRoot), h.handleUserRequest)
	log.Printf("initialized %s/user REST handler\n", apiRoot)
	// POST + PUT
	http.HandleFunc(fmt.Sprintf("%s/auction", apiRoot), h.handleAuctionRequest)
	log.Printf("initialized %s/auction REST handler\n", apiRoot)
	// POST
	http.HandleFunc(fmt.Sprintf("%s/auction/register", apiRoot), h.handleAuctionRegisterRequest)
	log.Printf("initialized %s/auction/register REST handler\n", apiRoot)
}

func (h handler) handleUserRequest(w http.ResponseWriter, r *http.Request) {
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
			fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
			return
		}

		response, err := h.logic.CreateUser(userRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
			return
		}
		json.NewEncoder(w).Encode(response)
	default:
		fmt.Fprintf(w, "{\"error\": \"Only POST methods are supported.\"}")
	}
}

func (h handler) handleAuctionRequest(w http.ResponseWriter, r *http.Request) {
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
			fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
			return
		}

		response, err := h.logic.CreateAuction(createAuctionRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
			return
		}
		json.NewEncoder(w).Encode(response)
	case "PUT":
		log.Println("handling PUT /api/auction")

		var startAuctionRequest model.StartAuctionRequest
		err := dec.Decode(&startAuctionRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
			return
		}

		response, err := h.logic.StartAuction(startAuctionRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
			return
		}
		json.NewEncoder(w).Encode(response)
	default:
		fmt.Fprintf(w, "{\"error\": \"Only POST and PUT methods are supported.\"}")
	}
}

func (h handler) handleAuctionRegisterRequest(w http.ResponseWriter, r *http.Request) {
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
			fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
			return
		}

		response, err := h.logic.RegisterAuction(auctionRegisterRequest)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
			return
		}
		json.NewEncoder(w).Encode(response)
	default:
		fmt.Fprintf(w, "{\"error\": \"Only POST methods are supported.\"}")
	}
}
