package main

import "github.com/gin-gonic/gin"

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	Router := gin.Default()

	Router.POST("/register", Register)
	Router.POST("/login", Login)
	Router.POST("/logout", Logout)

	protectedRouter := Router.Group("/")
	protectedRouter.Use(IsLogin())
	{
		protectedRouter.POST("/leaveM", LeaveMessage)
		protectedRouter.DELETE("/deleteM", DeleteMessage)
	}

	err = Router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
