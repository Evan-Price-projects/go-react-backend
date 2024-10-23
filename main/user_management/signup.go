package user_management

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SignupInfo struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Birthdate time.Time `json:"birthdate"`
}

func Signup(c *gin.Context) {
	var newSignupInfo SignupInfo
	if err := c.BindJSON(&newSignupInfo); err != nil {
		return
	}
	fmt.Print(newSignupInfo)
	c.IndentedJSON(http.StatusCreated, newSignupInfo)
}
