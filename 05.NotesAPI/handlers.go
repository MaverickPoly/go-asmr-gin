package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchAllNotes(c *gin.Context) {
	notes := make([]Note, 0)

	if err := DB.Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error fetching notes: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func CreateNote(c *gin.Context) {
	var note Note

	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Body!",
		})
		return
	}

	if note.Title == "" || note.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Some fields are missing!",
		})
		return
	}

	if err := DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error creating note: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, note)
}

func FetchNoteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid note id!",
		})
		return
	}

	var note Note
	if err := DB.Find(&note, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error fetching note with id %v!", id),
		})
		return
	}

	if note.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Note with id %v not found!", id),
		})
		return
	}

	c.JSON(http.StatusOK, note)
}

func DeleteNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid note id!",
		})
		return
	}

	var note Note
	if err := DB.Find(&note, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error fetching note with id %v!", id),
		})
		return
	}

	if note.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Note with id %v not found!", id),
		})
		return
	}

	if err := DB.Delete(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error deleting note with id %v: %v", id, err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, note)
}

func UpdateNote(c *gin.Context) {
	// Get ID Param
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid note id!",
		})
		return
	}

	// Find note with id
	var dbNote Note
	if err := DB.Find(&dbNote, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error fetching note with id %v!", id),
		})
		return
	}

	if dbNote.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Note with id %v not found!", id),
		})
		return
	}

	// Fetch Body
	var note Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Body!",
		})
		return
	}

	if err := DB.Model(dbNote).Updates(note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error updating todo with id %v: %v", id, err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, note)
}
