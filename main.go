package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "http://localhost:8081"

func main() {
	lc := loginClient(url+"/v1/login", "wings@gmail.com", "123456")
	log.Printf("Login response: %v", lc)

	Person := Person{
		Name:        "Maria",
		Age:         25,
		Email:       "felival@gmail.com",
		PhoneNumber: "999000666",
		Password:    "123456",
		Communities: []Community{
			{
				Name: "Community 1",
			},
			{
				Name: "Community 2",
			},
		},
	}

	fmt.Println("Created Person")
	createResponse := createPerson(url+"/v1/persons", lc.Data, &Person)
	fmt.Println(createResponse)

	fmt.Println("Updated Person")
	updatedResponse := updatePerson(url+"/v1/persons/", lc.Data, 11, &Person)
	fmt.Println(updatedResponse)

	fmt.Println("Delete by ID")
	genericResponse := deletePerson(url+"/v1/persons/", lc.Data, 12)
	fmt.Println(genericResponse)

	fmt.Println("Get All")
	getllResponse := getAllPerson(url+"/v1/persons", lc.Data)
	fmt.Println(getllResponse)

	fmt.Println("Get by ID")
	genericResponseGet := getByIDPerson(url+"/v1/persons", lc.Data, 13)
	fmt.Println(genericResponseGet)

}

func httpClient(method, url, token string, body io.Reader) (*http.Response, error) {
	var req *http.Request
	var err error
	if method == http.MethodGet || method == http.MethodDelete {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, body)
	}
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	return resp, nil
}
