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

	maxRand    = 100
	RandNumber = rand.Intn(maxRand)
	first      = 0
	last       = 100

	HTTPstatus = 0
	authStatus bool
)

func main() {

	route.Use(cors.Default())

	route.POST("/login", Login)
	route.GET("/authCheck", VerifyToken)

	/* Use Middleware */
	route.Use(MiddlewareAuthen())
	{
		route.GET("/key", func(c *gin.Context) {
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
		var tokenString = c.Request.Header.Get("Token")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
		} else {
			claims := jwt.MapClaims{}
			_token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("ACCESS_SECRET")), nil
			})
			if err != nil {
				c.AbortWithStatusJSON(401, gin.H{
					"message": "Unauthorized",
				})
				authStatus = false
			} else {
				authStatus = _token.Valid
				c.Next()
			}
		}
	}
}

func Guess(ctx *gin.Context) {
	var guessNumber, err = strconv.Atoi(ctx.Query("number"))
	if err != nil {
		ctx.JSON(203, gin.H{
			"message": "Not have parameter",
		})
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

func VerifyToken(ctx *gin.Context) {
	var tokenString = ctx.Request.Header.Get("Token")
	if tokenString == "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
	} else {
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})
		if err != nil {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
				"status":  false,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Authorized",
				"status":  token.Valid,
			})
		}
	}
}

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

/* test-user */
var user = User{
	ID:       1,
	Username: "test",
	Password: "testuser",
	Name:     "Test User",
}
