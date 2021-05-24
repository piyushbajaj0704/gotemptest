package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/pkg/errors"
)

// HomePage of the Fast
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Fast!")
	fmt.Println("Endpoint Hit: homePage")
}

// HandleRequests is the single place to track all endpoints
func handleRequests(r *mux.Router) {
	// replace http.HandleFunc with myRouter.HandleFunc
	r.HandleFunc("/", HomePage)

	// ------------------------------------------------------------------------------------------------------------------------
	// --------------------------------------------Endpoints for Organisation-----------------------------------------
	// ------------------------------------------------------------------------------------------------------------------------

	// swagger:operation GET /api/v1/orgs org getOrgs
	//
	// ---
	// summary: Return all the org details.
	// description: Orgs details will be returned else Error Not Found (404) will be returned.
	r.HandleFunc("/api/v1/orgs", CreateOrg).Methods("GET")
}

// CreateOrg to add a new org
func CreateOrg(w http.ResponseWriter, r *http.Request) {
	fmt.Print("hello")
	return
}
