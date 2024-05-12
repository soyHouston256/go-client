package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

func createPerson(url, token string, person *Person) GenericResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(person)
	if err != nil {
		log.Fatalf("Error encoding person data: %v", err)
	}

	response, err := httpClient(http.MethodPost, url, token, data)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusCreated {
		log.Fatalf("Error logging in %v", response.Status)
	}
	dataResponse := GenericResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}
	return dataResponse
}

func updatePerson(url, token string, id int, person *Person) GenericResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(person)
	if err != nil {
		log.Fatalf("Error encoding person data: %v", err)
	}

	response, err := httpClient(http.MethodPut, url+strconv.Itoa(id), token, data)
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
	dataResponse := GenericResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}
	return dataResponse
}

func deletePerson(url, token string, id int) GenericResponse {

	response, err := httpClient(http.MethodDelete, url+strconv.Itoa(id), token, nil)
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
	dataResponse := GenericResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}
	return dataResponse
}

func getAllPerson(url, token string) AllPersonsResponse {
	response, err := httpClient(http.MethodGet, url, token, nil)
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
	dataResponse := AllPersonsResponse{}
	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}
	return dataResponse
}

func getByIDPerson(url, token string, id int) GetResponse {
	response, err := httpClient(http.MethodGet, url+"/"+strconv.Itoa(id), token, nil)
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
	dataResponse := GetResponse{}
	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}
	return dataResponse
}
