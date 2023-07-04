package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// contact represents data about a contact.
type contact struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

// contacts slice to store contact data.
var contacts = []contact{
	{ID: "1", Name: "John Doe", PhoneNumber: "1234567890"},
	{ID: "2", Name: "Jane Smith", PhoneNumber: "9876543210"},
}

func main() {
	router := gin.Default()
	router.GET("/contacts", getContacts)
	router.GET("/contacts/:id", getContactByID)
	router.POST("/contacts", createContact)
	router.PUT("/contacts/:id", updateContact)
	router.DELETE("/contacts/:id", deleteContact)

	err := router.Run("localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
}

// getContacts responds with the list of all contacts as JSON.
func getContacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, contacts)
}

// getContactByID locates the contact whose ID value matches the id
// parameter sent by the client, then returns that contact as a response.
func getContactByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of contacts, looking for
	// a contact whose ID value matches the parameter.
	for _, contact := range contacts {
		if contact.ID == id {
			c.IndentedJSON(http.StatusOK, contact)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
}

// createContact adds a new contact with the given name and phone number.
func createContact(c *gin.Context) {
	var newContact contact

	// Call BindJSON to bind the received JSON to newContact.
	if err := c.BindJSON(&newContact); err != nil {
		log.Println("Error binding JSON:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	log.Println("Received data:", newContact)

	// Generate a new unique ID for the contact.
	newContact.ID = generateID()

	// Add the new contact to the slice.
	contacts = append(contacts, newContact)
	c.IndentedJSON(http.StatusCreated, newContact)
}

// updateContact updates the contact with the given ID.
func updateContact(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of contacts, looking for
	// the contact whose ID value matches the parameter.
	for i, contact := range contacts {
		if contact.ID == id {
			// Update the contact with the new values from JSON.
			if err := c.BindJSON(&contacts[i]); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
				return
			}

			c.IndentedJSON(http.StatusOK, contacts[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
}

// deleteContact deletes the contact with the given ID.
func deleteContact(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of contacts, looking for
	// the contact whose ID value matches the parameter.
	for i, contact := range contacts {
		if contact.ID == id {
			// Delete the contact from the slice.
			contacts = append(contacts[:i], contacts[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
}

// generateID generates a new unique ID for a contact.
func generateID() string {
	// In a real-world scenario, you would use a more robust
	// method to generate unique IDs, such as UUIDs.
	// For simplicity, a sequential ID is used here.
	return strconv.Itoa(len(contacts) + 1)
}
