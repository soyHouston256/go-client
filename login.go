package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func loginClient(url, email, password string) LoginResponse {
	login := Login{
		Email:    email,
		Password: password,
	}
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(&login)
	if err != nil {
		log.Fatalf("Error encoding login data: %v", err)
	}
	response, err := httpClient(http.MethodPost, url, "", data)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Error logging in %v", response.Status)
	}
	dataResponse := LoginResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}
	return dataResponse
}
