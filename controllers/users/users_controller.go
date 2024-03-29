package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sharkx018/bookstore_oauth-go/oauth"
	"github.com/sharkx018/bookstore_utils-go/rest_errors"

	"github.com/sharkx018/bookstore_users-api/domain/users"
	"github.com/sharkx018/bookstore_users-api/services"
	"net/http"
	"strconv"
)

func getUserId(userID string) (int64, rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(userID, 10, 64)
	if userErr != nil {
		err := rest_errors.NewBadRequestError("user id should be a number")
		return -1, err
	}
	return userId, nil
}

func CreateUser(c *gin.Context) {
	var user users.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	fmt.Println(user)

	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	isPublic := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusCreated, result.Marshal(isPublic))

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func GetUser(c *gin.Context) {

	if err := oauth.AuthenticationRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
	}

	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	fmt.Println("===>>>>>> GetCallerId: ", oauth.GetCallerId(c.Request))
	fmt.Println("===>>>>>> UserID: ", user.Id)
	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshal(false))
		return
	}

	isPublic := oauth.IsPublic(c.Request)
	c.JSON(http.StatusOK, user.Marshal(isPublic))
}

func UpdateUser(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
	}

	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UserService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
		return
	}
	isPublic := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusOK, result.Marshal(isPublic))

}

func DeleteUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
	}

	if err := services.UserService.DeleteUser(userId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	isPublic := c.GetHeader("X-Public") == "true"

	results := users.Marshal(isPublic)

	c.JSON(http.StatusOK, results)
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	user, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, user)
}
