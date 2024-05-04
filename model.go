package main

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Person struct {
	Name        string
	Age         int
	Email       string
	PhoneNumber string
	Password    string
	Communities []Community
}

type Community struct {
	Name     string
	PersonID uint
}

type Persons []Person

type GenericResponse struct {
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
}

type LoginResponse struct {
	GenericResponse
	Data string `json:"data"`
}

type AllPersonsResponse struct {
	GenericResponse
	Data []Person `json:"data"`
}

type GetResponse struct {
	GenericResponse
	Data Person `json:"data"`
}
