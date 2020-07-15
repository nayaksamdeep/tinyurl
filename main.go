/*
 * This program measures the number of glasses of water a person drinks a day
 * It will an user to input the number of glasses he has drank
 * It will also show him the total number of cups he has consumed for the day
 */

package main

import  (
     "fmt"
//     "log"
//     "github.com/nayaksamdeep/tinyurl/Config"
//     "github.com/nayaksamdeep/contact-list-service/Models"
     "github.com/nayaksamdeep/tinyurl/Routes"
//     "github.com/jinzhu/gorm"
     _"github.com/mattn/go-sqlite3"
)

var err error

func main() {

/*
	// Creating a connection to the database

	Config.DB, err = gorm.Open("sqlite3", "./gorm.db")

	if err != nil {
                log.Fatal(err)
	}

	defer Config.DB.Close()

	// run the migrations: todo struct
	Config.DB.AutoMigrate(&Models.Contact{})
*/

        fmt.Println("hi there! Welcome to Contact List Service\n");

        r := Routes.SetupRouter()
        // running
        r.Run() 
}
