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
	route = gin.Default()

	maxRand    = 5
	RandNumber = rand.Intn(maxRand)
	first      = 0
	last       = 100

	HTTPstatus = 0

	token = 1
)

func main() {

	route.Use(cors.Default())

	route.POST("/login", Login)
	route.POST("/authCheck", authenCheck)
	route.GET("/key", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"number": RandNumber,
		})
	})

	/* Use Middleware */
	route.Use(MiddlewareAuthen())
	{
		route.GET("/middle", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"number": RandNumber,
			})
		})
		route.GET("/guess", Guess)
	}

	log.Fatal(route.Run(":8080"))
}

/* Middleware Authentication Check */
func MiddlewareAuthen() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token == 0 {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func Guess(ctx *gin.Context) {
	var guessNumber, err = strconv.Atoi(ctx.Query("number"))
	if err != nil {
		return
	}
	if RandNumber == guessNumber {
		HTTPstatus = 201
		ctx.JSON(HTTPstatus, gin.H{
			"number":  guessNumber,
			"status":  true,
			"message": "Congratulations!",
		})
		RandNumber = rand.Intn(maxRand)
	} else {
		HTTPstatus = 202
		var message = ""
		if guessNumber < RandNumber {
			message = "is too low"
			first = guessNumber
		} else {
			if last > maxRand {
				last = maxRand
			} else {
				last = guessNumber
			}
			message = "is too high"
		}
		ctx.JSON(HTTPstatus, gin.H{
			"number":  guessNumber,
			"status":  false,
			"first":   first,
			"last":    last,
			"message": message,
		})
	}
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

func authenCheck(ctx *gin.Context) {
	tokenString := ctx.Query("jwt_token")
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, token.Valid)
}
