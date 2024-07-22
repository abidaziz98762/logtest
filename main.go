package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", homeHandler)
	r.GET("/about", aboutHandler)
	r.GET("/contact", contactHandler)

	r.Run(":9999")
}

func homeHandler(c *gin.Context) {
	c.String(200, "Welcome to the home page!")
}

func aboutHandler(c *gin.Context) {
	c.String(200, "This is the about page.")
}

func contactHandler(c *gin.Context) {
	c.String(200, "Contact us at: example@email.com")
}
