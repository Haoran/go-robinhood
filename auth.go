package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/go-resty/resty"
	"fmt"
)

type Auth struct {
	Token string `json:"token"`
}

type Profile struct {
}

type LoginData struct {
	Username        string   `json:"username,omitempty"`
	Password	string   `json:"password,omitempty"`
}

func CreateAuthEndpoint(w http.ResponseWriter, req *http.Request) {
	var loginData LoginData
	_ = json.NewDecoder(req.Body).Decode(&loginData)

	fmt.Print(loginData)

	client1 := resty.New()
	resp, err := client1.R().
		SetHeader("Accept", "application/json").
		SetBody(loginData).
		Post("https://api.robinhood.com/api-token-auth/")

	auth := &Auth{
	}

	error := json.Unmarshal(resp.Body(), auth)

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nError: %v", error)
	fmt.Printf("\nResponse: %v", resp)
	fmt.Printf("\nInfo: %v", auth)

	json.NewEncoder(w).Encode(auth)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user/auth", CreateAuthEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":12311", router))
}