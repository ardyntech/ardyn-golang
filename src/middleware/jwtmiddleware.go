package middleware

//-------------------------------------------------------------

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"ardyngolang/src/certsec"
	"ardyngolang/src/responses"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//-------------------------------------------------------------

// http://www.inanzzz.com/index.php/post/kdl9/creating-and-validating-a-jwt-rsa-token-in-golang

func validateToken(encodedToken string) (*jwt.Token, error) {

	key, err := jwt.ParseRSAPublicKeyFromPEM(certsec.GetPublicKey())

	if err != nil {

		return nil, fmt.Errorf("Validate: parse key: %w", err)

	}

	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {

		if _, isvalid := token.Method.(*jwt.SigningMethodRSA); !isvalid {

			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])

		}

		return key, nil

	})
}

//-------------------------------------------------------------

// Custom Gin Middleware to check the validity of the bearer token
// first before proceeding with any request.

func AuthorizeJWT() gin.HandlerFunc {

	var response responses.JsonResponse

	return func(c *gin.Context) {

		const BEARER_SCHEMA = "Bearer"

		authHeader := c.GetHeader("Authorization")

		if len(authHeader) < 1 {

			log.Println("No AUTH header found.")

			response.Status = "Error"

			response.Message = "No AUTH header found."

			//c.AbortWithStatus(http.StatusUnauthorized)

			c.JSON(http.StatusUnauthorized, response)

			c.Abort()

			return

		}

		tokenString := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])

		// Does the bearer token exist?
		if len(tokenString) == 0 {

			//c.AbortWithStatus(http.StatusUnauthorized)

			response.Status = "Error"

			response.Message = "No bearer token found."

			c.JSON(http.StatusUnauthorized, response)

			c.Abort()

			return

		}

		// All ok. Now validate the token
		token, err := validateToken(tokenString)

		if err != nil {

			//c.AbortWithStatus(http.StatusUnauthorized)

			log.Println("Token Error: ", err)

			response.Status = "Error"

			response.Message = "Invalid or expired token found. Please try signing in again."

			response.Error = err.Error()

			c.JSON(http.StatusUnauthorized, response)

			c.Abort()

			return

		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {

			c.AbortWithStatus(http.StatusUnauthorized)

			log.Println("Claims Error: ", err)

			c.Abort()

			return

		}

		// Finally, get the user id from the token and set it
		// for all future requests.
		// If you are using ArdynCore, then the userId is stored
		// under "userId" in the claims section of the jwt token.
		c.Set("user_id", claims["userId"])

		c.Next()

	}

}

//-------------------------------------------------------------
