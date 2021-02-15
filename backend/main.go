package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	route      = gin.Default()
	RandNumber = rand.Intn(100)
	status     = 200
)

func main() {

	route.Use(cors.Default())

	route.POST("/login", Login)
	route.GET("/key", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"number": RandNumber,
		})
	})
	route.GET("/guess", Guess)

	log.Fatal(route.Run(":8080"))
}

func Guess(ctx *gin.Context) {
	var guessNumber, err = strconv.Atoi(ctx.Query("guessNumber"))
	if err != nil {
		return
	}
	if RandNumber == guessNumber {
		status = 201
		RandNumber = rand.Intn(5)
	} else {
		status = 202
	}
	ctx.JSON(status, gin.H{
		"guessNumber": guessNumber,
	})
}

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

var user = User{
	ID:       1,
	Username: "test",
	Password: "testuser",
	Name:     "Test User",
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func CreateToken(userId uint64) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
