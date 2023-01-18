package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// const dbFile string = "urls.db"

// // album represents data about a record album.
// type url struct {
// 	Title  string  `json:"title"`
// 	Short string  `json:"short"`
// 	Original  string `json:"original"`
// }

// // albums slice to seed record album data.
// var urls = []url{}

// func getShortedLinks(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, urls)
// }

// func saveShortedLink(c *gin.Context) {
// 	var newUrl url

// 	if err := c.BindJSON(&newUrl); err != nil {
// 		return
// 	}
// 	// Add the new album to the slice.
// 	urls = append(urls, newUrl)
// 	c.IndentedJSON(http.StatusCreated, urls)
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/", getShortedLinks)
// 	router.POST("/", saveShortedLink)

// 	router.Run("localhost:8080")
// }
