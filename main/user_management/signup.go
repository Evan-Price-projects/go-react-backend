package user_management

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	SignIn    SignIn `json:"signIn"`
	Email     string `json:"email"`
}

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {
	var newSignupInfo User
	if err := c.BindJSON(&newSignupInfo); err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(newSignupInfo)
	c.IndentedJSON(http.StatusCreated, newSignupInfo)
}
