package app

import (
	"github.com/aftaab60/store_users-api/controllers/ping"
	"github.com/aftaab60/store_users-api/controllers/user"
)

func mapUrl() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", user.GetUser)
	router.POST("/users", user.CreateUser)
	router.PUT("/users/:user_id", user.UpdateUser)
	router.PATCH("/users/:user_id", user.UpdateUser)
}
