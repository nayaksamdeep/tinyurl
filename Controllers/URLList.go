package Controllers

import (
	"net/http"
        "fmt"
        "github.com/nayaksamdeep/tinyurl/Models"
	"github.com/gin-gonic/gin"
)

func GetHomePage (c *gin.Context) {
       c.HTML(
             http.StatusOK,
             "index.html",
             gin.H{
                 "title": "Home Page",
             },
       )
}

func GetAllUrlInfo(c *gin.Context) {

	var urlstruct []Models.RedirectUrl
	err := Models.GetAllUrl(&urlstruct)
	if err != nil {
                fmt.Println("Request aborted with Status Not Found")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
                fmt.Println("Request processed with Status OK")
		c.JSON(http.StatusOK, urlstruct)
	}
}

func ConvertAUrl(c *gin.Context) {
	var urlstruct Models.RedirectUrl
	c.BindJSON(&urlstruct)
	err := Models.ConvertAUrl(&urlstruct)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
		c.JSON(http.StatusOK, urlstruct)
                fmt.Println("Request processed with Status OK")
	}
}

func RedirectAUrl(c *gin.Context) {
	var urlstruct Models.RedirectUrl
	id := c.Params.ByName("id")
	err := Models.RedirectAUrl(&urlstruct, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
                fmt.Println("Request processed with Status OK")
		c.JSON(http.StatusOK, urlstruct)
		// c.Redirect(http.StatusPermanentRedirect, urlstruct.url)
	}
}

func UpdateAUrl(c *gin.Context) {
	var urlstruct Models.RedirectUrl
	id := c.Params.ByName("id")
	err := Models.UpdateAUrl(&urlstruct, id)
	if err != nil {
		c.JSON(http.StatusNotFound, urlstruct)
	}
	c.BindJSON(&urlstruct)
	err = Models.UpdateAUrl(&urlstruct, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
		c.JSON(http.StatusOK, urlstruct)
                fmt.Println("Request processed with Status OK")
	}
}

func DeleteAUrl(c *gin.Context) {
	var urlstruct Models.RedirectUrl
	id := c.Params.ByName("id")
	err := Models.DeleteAUrl(&urlstruct, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
                fmt.Println("Request processed with Status OK")
	}
}
