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
		v1.GET("tinyurl/:id", Controllers.RedirectAUrl)
		v1.PUT("tinyurl/:id", Controllers.UpdateAUrl)
		v1.DELETE("tinyurl/:id", Controllers.DeleteAUrl)
	}

	return r
}
