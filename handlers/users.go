package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ArinaPereteatcu26/Pentagon/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var user db.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "failed to parse body")
		return
	}
	if len(user.Login) < 3 || len(user.Password) < 6 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Login should be at least 3 chars. Password should be at least 6.")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	user.Hash = string(hash)
	_, err = db.AddUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
}

func LogIn(c *gin.Context) {
	var user db.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "failed to parse body")
		return
	}

	dbUser, err := db.GetUser(user.Login)
	fmt.Println(dbUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Hash), []byte(user.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims =
		jwt.StandardClaims{
			Subject:   strconv.Itoa(dbUser.PersonId),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		}
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, tokenString)
}
