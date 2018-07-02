package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)

		if loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
