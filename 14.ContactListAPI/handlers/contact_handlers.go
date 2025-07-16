package handlers

import (
	"fmt"
	"net/http"

	"14.ContactListAPI/config"
	"14.ContactListAPI/models"
	"14.ContactListAPI/utils"
	"github.com/gin-gonic/gin"
)

func FetchAllContacts(c *gin.Context) {
	contacts := make([]models.Contact, 0)

	if err := config.DB.Find(&contacts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch all contacts: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "Fetched contacts successfully!",
		"data": contacts,
	})
}

func CreateContact(c *gin.Context) {
	var contact models.Contact

	if err := c.ShouldBindBodyWithJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid body!",
		})
		return
	}

	if contact.FirstName == "" || contact.SecondName == "" || contact.Email == "" || contact.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Some fields are missing!",
		})
		return
	}

	if err := config.DB.Create(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to create contact: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "Created contact successfully!",
		"data": contact,
	})
}

func FetchContact(c *gin.Context) {
	var contact models.Contact
	status, id, err := utils.GetContactById(c, &contact)

	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"msg":  fmt.Sprintf("Contact with id %v fetched successfully!", id),
		"data": contact,
	})
}

func DeleteContact(c *gin.Context) {
	var contact models.Contact
	status, id, err := utils.GetContactById(c, &contact)

	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := config.DB.Delete(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to delete contact with id %v: %v", id, err.Error()),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("Contact with id %v deleted successfully!", id),
		"data": contact,
	})
}

func UpdateContact(c *gin.Context) {
	var dbContact models.Contact
	status, id, err := utils.GetContactById(c, &dbContact)

	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	var contact models.Contact
	if err := c.ShouldBindBodyWithJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Body!",
		})
		return
	}

	if err := config.DB.Model(&dbContact).Updates(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to update contact with id %v: %v", id, err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("Contact with id %v updated successfully!", id),
		"data": dbContact,
	})
}
