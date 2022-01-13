// Our rest API package

package main

import (
	"ardyngolang/src/env"
	"ardyngolang/src/middleware"
	"ardyngolang/src/responses"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//--------------------------------------------------------------------------

func startHttp() {

	// New Gin instance
	r := gin.New()

	// Use our custom DB logger
	//r.Use(middleware.LogToDB())

	// This route is UNPROTECTED and does not require a jwt token
	// to be in the authorization header
	r.GET("/v1/test/", testEndpoint)

	// This route is PROTECTED and requires a valid jwt token to be
	// in the authorization header
	r.GET("/v1/test/me", middleware.AuthorizeJWT(), testMeEndpoint)

	// Start the HTTP server
	r.Run(":" + strconv.Itoa(env.Config.Port))

}

//--------------------------------------------------------------------------

func testEndpoint(c *gin.Context) {

	var response responses.JsonResponse

	response.Status = "Ok!"
	response.Message = "Unprotected endpoint is working!"

	c.JSON(http.StatusOK, response)

}

//--------------------------------------------------------------------------

func testMeEndpoint(c *gin.Context) {

	var response responses.JsonResponse

	// This "user_id" has been put in for us by the jwtmiddleware
	userId, valid := c.Get("user_id")

	if !valid {

		response.Status = "Error"
		response.Message = "Couldn't find the User ID, or none was provided in the token."

		c.JSON(http.StatusUnauthorized, response)

		return

	}

	uid := fmt.Sprintf("%v", userId)

	response.Status = "Ok!"
	response.Message = "Protected endpoint is working! User ID: " + uid

	c.JSON(http.StatusOK, response)

}

//--------------------------------------------------------------------------
