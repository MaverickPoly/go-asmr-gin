package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"14.ContactListAPI/config"
	"14.ContactListAPI/models"
	"github.com/gin-gonic/gin"
)

func GetContactById(c *gin.Context, contact *models.Contact) (int, int, error) {
	id, err := strconv.Atoi(c.Param("contactId"))

	if err != nil {
		return http.StatusBadRequest, id, errors.New("invalid contact id")
	}

	if err := config.DB.Find(contact, id).Error; err != nil {
		return http.StatusInternalServerError, id, fmt.Errorf("failed to create contact: %v", err.Error())
	}

	if contact.ID == 0 {
		return http.StatusNotFound, id, fmt.Errorf("contact with id %v not found", id)
	}

	return http.StatusOK, id, nil
}
