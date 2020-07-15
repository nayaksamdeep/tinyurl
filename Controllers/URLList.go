package Controllers

import (
//	"net/http"
//        "fmt"
//        "github.com/nayaksamdeep/tinyurl/Models"
	"github.com/gin-gonic/gin"
)

func GetAllUrlInfo(c *gin.Context) {
/*
	var contact []Models.Contact
	err := Models.GetAllContacts(&contact)
	if err != nil {
                fmt.Println("Request aborted with Status Not Found")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
                fmt.Println("Request processed with Status OK")
		c.JSON(http.StatusOK, contact)
	}
*/
}

func ConvertAUrl(c *gin.Context) {
/*
	var contact Models.Contact
	c.BindJSON(&contact)
	err := Models.CreateAContact(&contact)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
		c.JSON(http.StatusOK, contact)
                fmt.Println("Request processed with Status OK")
	}
*/
}

func RedirectAUrl(c *gin.Context) {
/*
	id := c.Params.ByName("id")
	var contact Models.Contact
	err := Models.GetAContact(&contact, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
                fmt.Println("Request processed with Status OK")
		c.JSON(http.StatusOK, contact)
	}
*/
}

func UpdateAUrl(c *gin.Context) {
/*
	var contact Models.Contact
	id := c.Params.ByName("id")
	err := Models.GetAContact(&contact, id)
	if err != nil {
		c.JSON(http.StatusNotFound, contact)
	}
	c.BindJSON(&contact)
	err = Models.UpdateAContact(&contact, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
		c.JSON(http.StatusOK, contact)
                fmt.Println("Request processed with Status OK")
	}
*/
}

func DeleteAUrl(c *gin.Context) {
/*
	var contact Models.Contact
	id := c.Params.ByName("id")
	err := Models.DeleteAContact(&contact, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
                fmt.Println("Request processed with Status OK")
	}
*/
}
