/*
 * This program measures the number of glasses of water a person drinks a day
 * It will an user to input the number of glasses he has drank
 * It will also show him the total number of cups he has consumed for the day
 */

package main

import  (
     "fmt"
     "log"
     "github.com/nayaksamdeep/tinyurl/Models"
     "github.com/nayaksamdeep/tinyurl/Routes"
     "github.com/jinzhu/gorm"
     _"github.com/mattn/go-sqlite3"
     "github.com/nayaksamdeep/tinyurl/docs"
     "github.com/gin-gonic/gin"
     swaggerFiles "github.com/swaggo/files" // swagger embed files
     ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var err error

// @title Tinyservice Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email kannanbalu7@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /api/v1
func main() {

        docs.SwaggerInfo.Title = "TinyService API";
        docs.SwaggerInfo.Description= "TinyService API - Samdeep/Kannan";
        docs.SwaggerInfo.Version = "1.0";
        docs.SwaggerInfo.BasePath = "/v1"; 
        docs.SwaggerInfo.Host = "localhost:7000";
        docs.SwaggerInfo.Schemes = []string{"http"};

	// Creating a connection to the database

	Models.DB, err = gorm.Open("sqlite3", "./tinyurlgorm.db")

	if err != nil {
                log.Fatal(err)
	}

	defer Models.DB.Close()

	// run the migrations: todo struct
	Models.DB.AutoMigrate(&Models.RedirectUrl{})

        fmt.Println("hi there! Welcome to Contact List Service\n");

       
        server := gin.Default()
        server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)); 
        //Need to spawn it on a separate thread as this is a synchronous and blocking all
        go server.Run(":5000");

        r := Routes.SetupRouter()
        r.Run() 
}
