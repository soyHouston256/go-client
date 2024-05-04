package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
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

func updatePerson(url, id, token string, person *Person) GenericResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(person)
	if err != nil {
		log.Fatalf("Error encoding person data: %v", err)
	}

	response, err := httpClient(http.MethodPut, url+id, token, data)
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

func deletePerson(url, id, token string, person *Person) GenericResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(person)
	if err != nil {
		log.Fatalf("Error encoding person data: %v", err)
	}

	response, err := httpClient(http.MethodDelete, url+id, token, data)
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

func getAllPerson(url, token string, person *Person) AllPersonsResponse {
	//data := bytes.NewBuffer([]byte{})
	//err := json.NewEncoder(data).Encode(person)
	//if err != nil {
	//	log.Fatalf("Error encoding person data: %v", err)
	//}

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

func getPerson(url, id, token string, person *Person) GetResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(person)
	if err != nil {
		log.Fatalf("Error encoding person data: %v", err)
	}

	response, err := httpClient(http.MethodGet, url+id, token, nil)

	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Error logging in %v", response.Status)
	}
	dataResponse := GetResponse{}
	err = json.Unmarshal(responseBody, &dataResponse)
	if err != nil {
		log.Fatalf("Error decoding response body: %v", err)
	}
	return dataResponse
}

func httpClientGet(url, token string) AllPersonsResponse {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", token)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	respBody, _ := io.ReadAll(response.Body)
	dataResponse := AllPersonsResponse{}
	err = json.Unmarshal(respBody, &dataResponse)
	if err != nil {
		log.Fatal(err)
	}
	return dataResponse
}
