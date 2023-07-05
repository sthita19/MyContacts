package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

var contacts []Contact

func main() {
	router := gin.Default()

	// Get all contacts
	router.GET("/contacts", func(c *gin.Context) {
		c.JSON(http.StatusOK, contacts)
	})

	// Get a single contact
	router.GET("/contacts/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, contact := range contacts {
			if contact.ID == id {
				c.JSON(http.StatusOK, contact)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
	})

	// Create a contact
	router.POST("/contacts", func(c *gin.Context) {
		var contact Contact

		if err := c.ShouldBindJSON(&contact); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		contacts = append(contacts, contact)
		c.Status(http.StatusCreated)
	})

	// Update a contact
	router.PUT("/contacts/:id", func(c *gin.Context) {
		id := c.Param("id")

		var updatedContact Contact

		if err := c.ShouldBindJSON(&updatedContact); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, contact := range contacts {
			if contact.ID == id {
				contacts[i] = updatedContact
				c.Status(http.StatusOK)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
	})

	// Delete a contact
	router.DELETE("/contacts/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i, contact := range contacts {
			if contact.ID == id {
				contacts = append(contacts[:i], contacts[i+1:]...)
				c.Status(http.StatusOK)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
	})

	router.Run(":8000")
}
