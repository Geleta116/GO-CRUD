package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)


type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}


var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}



func postAlbums(c *gin.Context){
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	  }
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}



func getAlbumByID(c *gin.Context){
	id := c.Param("id");
	for _, a := range albums{
		if id == a.ID {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound,"Album doesn't exist" )
	
}



func deleteAlbum(c *gin.Context){
	id := c.Param("id")
	for index, a := range albums{
		if id == a.ID {
			albums = append(albums[:index], albums[index+ 1:]...);
			c.IndentedJSON(http.StatusNoContent, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, "Album not found")

}



func updateAlbum(c *gin.Context){
	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil{
		c.IndentedJSON(http.StatusBadRequest, "failed to update user")
		return

	}

	id := c.Param("id")
	for index, alb := range albums{
		if id == alb.ID{
			albums[index] = updatedAlbum
			c.IndentedJSON(http.StatusCreated, alb)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest,"no such album")

}




func main(){
	 
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbum)
	router.PUT("/albums/:id", updateAlbum)
	
	router.Run("localhost:8080")
}











