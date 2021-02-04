package main

import "fmt"

type User struct {
	id   int
	name string
	Address
}

type Tech struct {
	id   int
	tech []TechDetails
}

type Contact struct {
	id int
	ContactDetails
}

type Address struct {
	area    string
	country string
}

type TechDetails struct {
	tech string
	exp  float64
}
type ContactDetails struct {
	email string
	phone string
}
type MergeUser struct {
	id   int
	name string
	Address
	techdets []TechDetails
	email    string
	phone    string
}

func merge(user []User, tech []Tech, contact []Contact) {
	countryCode := map[string]string{
		"IND": "+91",
		"UK":  "+41",
	}
	for i := 0; i < len(user); i++ {
		if user[i].country == "IND" {
			getPhone := contact[i].phone
			contact[i].phone = countryCode["IND"] + getPhone
		} else {
			getPhone := contact[i].phone
			contact[i].phone = countryCode["UK"] + getPhone
		}

	}
	store := make([]MergeUser, len(user))
	for index, user := range user {

		mergeStructVariable := MergeUser{}

		//Entering values of user

		mergeStructVariable.id = user.id
		mergeStructVariable.name = user.name
		mergeStructVariable.Address = user.Address

		for _, value := range tech {
			if value.id == user.id {
				mergeStructVariable.techdets = value.tech
			}
		}

		for _, value := range contact {
			if value.id == user.id {
				mergeStructVariable.email = value.email
				mergeStructVariable.phone = value.phone
			}
		}
		store[index] = mergeStructVariable
	}
	fmt.Println(store)
}
func main() {
	user1 := User{1, "Radha", Address{"Bopal", "IND"}}
	tech1 := Tech{1, []TechDetails{{"react", 3}, {"Golang", 1.5}}}
	contact1 := Contact{1, ContactDetails{"radha.kotech@bacancy.com", "123457890"}}
	users := []User{user1}
	tech := []Tech{tech1}
	contact := []Contact{contact1}
	merge(users, tech, contact)
}
