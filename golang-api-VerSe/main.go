package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Contact represents the structure of a contact
type Contact struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	ProfileImage string `json:"profile_image"`
}

var contacts []Contact

// GetContactsHandler retrieves all contacts
func GetContactsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contacts)
}

// GetContactHandler retrieves a specific contact by ID
func GetContactHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, contact := range contacts {
		if contact.ID == params["id"] {
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

// CreateContactHandler creates a new contact
func CreateContactHandler(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contact.ID = generateID()
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

// UpdateContactHandler updates an existing contact
func UpdateContactHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, contact := range contacts {
		if contact.ID == params["id"] {
			var updatedContact Contact
			_ = json.NewDecoder(r.Body).Decode(&updatedContact)
			updatedContact.ID = contact.ID
			contacts[i] = updatedContact
			json.NewEncoder(w).Encode(updatedContact)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

// DeleteContactHandler deletes a contact
func DeleteContactHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, contact := range contacts {
		if contact.ID == params["id"] {
			contacts = append(contacts[:i], contacts[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

func generateID() string {
	// Generate a unique ID for a contact (you can use your preferred ID generation method)
	// This is a simple example that uses the length of the contacts slice as the ID
	return string(len(contacts) + 1)
}

func main() {
	router := mux.NewRouter()

	contacts = append(contacts, Contact{
		ID:           generateID(),
		Name:         "John Doe",
		PhoneNumber:  "1234567890",
		ProfileImage: "profile.jpg",
	})

	router.HandleFunc("/contacts", GetContactsHandler).Methods("GET")
	router.HandleFunc("/contacts/{id}", GetContactHandler).Methods("GET")
	router.HandleFunc("/contacts", CreateContactHandler).Methods("POST")
	router.HandleFunc("/contacts/{id}", UpdateContactHandler).Methods("PUT")
	router.HandleFunc("/contacts/{id}", DeleteContactHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
