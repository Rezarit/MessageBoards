package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LeaveMessage(client *gin.Context) {
	var message Messages
	if err := client.ShouldBind(&message); err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := "INSERT INTO message(content) VALUES (?);"
	result, err := db.Exec(cmd, message.Content)
	ErrorDetection(client, err)

	_, err = result.LastInsertId()
	ErrorDetection(client, err)
	client.JSON(http.StatusOK, gin.H{"message": "Successfully left messages"})
}

func DeleteMessage(client *gin.Context) {
	var message Messages
	if err := client.ShouldBind(&message); err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := "DELETE FROM message WHERE user_id = ?;"
	result, err := db.Exec(cmd, message.UserId)
	ErrorDetection(client, err)

	client.JSON(http.StatusOK, gin.H{"message": "Successfully delete messages"})

	_, err = result.RowsAffected()
	ErrorDetection(client, err)
}
