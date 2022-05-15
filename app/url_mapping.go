package app

import (
	"github.com/sharkx018/bookstore_users-api/controllers/ping"
	"github.com/sharkx018/bookstore_users-api/controllers/users"
)

func mapUrls() {

	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
	router.GET("/internal/users/search", users.Search)

}
