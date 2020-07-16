package Controllers

import (
	"net/http"
        "fmt"
        "github.com/nayaksamdeep/tinyurl/Models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
        fmt.Println("params: ", c.Keys, c.FullPath());
	var urlstruct Models.RedirectUrl
	//val := c.BindJSON(&urlstruct)
        val := c.ShouldBindWith(&urlstruct, binding.FormPost);
        fmt.Println("Binding: ", val);
	err := Models.ConvertAUrl(&urlstruct)
	if err != nil {
                fmt.Println("Error: " , err);
		c.AbortWithStatus(http.StatusNotFound)
                fmt.Println("Request aborted with Status Not Found")
	} else {
                var id = fmt.Sprint(urlstruct.ID);
                urlstr := "/v1?url=" +  urlstruct.Url + "&tinyurl=Here is your shortened URL: <font color=\"green\"> http://" + id  + "/" + urlstruct.Url + " </font>";
                c.Redirect(302, urlstr);
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
