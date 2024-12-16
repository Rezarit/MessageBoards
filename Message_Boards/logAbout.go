package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(client *gin.Context) {
	var user Users
	if err := client.ShouldBind(&user); err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if 3 > len(user.Nickname) {
		client.JSON(http.StatusBadRequest, gin.H{"error": "Nickname is too short"})
		return
	}
	if 3 > len(user.Password) {
		client.JSON(http.StatusBadRequest, gin.H{"error": "Password is too short"})
		return
	}

	cmd := "INSERT INTO users(nickname,username,password) VALUES (?,?,?);"
	result, err := db.Exec(cmd, user.Nickname, user.Username, user.Password)
	if err != nil {
		client.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = result.RowsAffected()
	if err != nil {
		client.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func Login(client *gin.Context) {
	var user Users
	if err := client.ShouldBind(&user); err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := client.Query("password")
	if password != user.Password {
		client.JSON(http.StatusBadRequest, gin.H{"error": "Password does not match"})
	}

	client.SetCookie("islogin", "true", 600, "/", "", false, true)
	client.JSON(http.StatusOK, gin.H{"message": "User logged in"})
}

func IsLogin() gin.HandlerFunc {
	return func(client *gin.Context) {
		cookieValue, err := client.Cookie("islogin")
		if err != nil || cookieValue != "true" {
			client.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			client.Abort()
			return
		}
		client.Next()
	}
}

func Logout(client *gin.Context) {
	client.SetCookie("islogin", "false", -1, "/", "", false, true)
	client.JSON(http.StatusOK, gin.H{"message": "User logged out"})
}
