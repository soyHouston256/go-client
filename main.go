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

	//Person := Person{
	//	Name:        "Maria",
	//	Age:         25,
	//	Email:       "maria@gmail.com",
	//	PhoneNumber: "1234567890",
	//	Password:    "123456",
	//	Communities: []Community{
	//		{
	//			Name: "Community 1",
	//		},
	//		{
	//			Name: "Community 2",
	//		},
	//	},
	//}

	//genericResponse := createPerson(url+"/v1/persons", lc.Data, &Person)
	//fmt.Println(genericResponse)

	//genericResponse := updatePerson(url+"/v1/persons/", "12", lc.Data, &Person)
	//fmt.Println(genericResponse)
	//
	//genericResponse = deletePerson(url+"/v1/persons/", "6", lc.Data, &Person)
	//fmt.Println(genericResponse)

	//genericResponse := getAllPerson(url+"/v1/persons", lc.Data, nil)
	//fmt.Println(genericResponse)
	//genericResp := httpClientGet(url+"/v1/persons", lc.Data)
	//fmt.Println(genericResp)

	genericResponse := getAllPerson(url+"/v1/persons", lc.Data, nil)
	fmt.Println(genericResponse)

}

func httpClient(method, url, token string, body io.Reader) (*http.Response, error) {
	var req *http.Request
	var err error
	if method == http.MethodGet {
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
