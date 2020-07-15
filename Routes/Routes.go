package Routes

import (
        "github.com/nayaksamdeep/tinyurl/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
        r.LoadHTMLGlob("templates/*");

	v1 := r.Group("/v1")
	{
		v1.GET("/", Controllers.GetHomePage)
		v1.GET("ListURL", Controllers.GetAllUrlInfo)
		v1.POST("ConvertURL", Controllers.ConvertAUrl)
		v1.GET("ListURL/:id", Controllers.RedirectAUrl)
		v1.PUT("ListURL/:id", Controllers.UpdateAUrl)
		v1.DELETE("ListURL/:id", Controllers.DeleteAUrl)
	}

	return r
}
