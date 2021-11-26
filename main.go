package main

import (
	// Go importing pkg primary
	// "fmt"
	"net/http"
	"strconv"

	// Go importing module
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	
	// Go importing owner
	"github.com/huanNguyen97/huan-sa-hsu-golang/layer/B_BUS_layer"
	"github.com/huanNguyen97/huan-sa-hsu-golang/DTO"
)

func main() {
	// APP CREATED 
	// --------------------------------------------------------------
	// --------------------------------------------------------------
	app := gin.Default()

	// Enable CORS to allow all origin.
	app.Use(cors.Default())

	// Read games
	app.GET("/", func(c *gin.Context) {
		games := B_BUS_layer.ReadGamesBUS()
		c.JSON(http.StatusOK, gin.H{"game_list": games})
	})
	// End read
	
	// Create game
	app.POST("/create-game", func(c *gin.Context) {
		game_created := new(DTO.Game)

		game_created.ID	 		= ""
		game_created.Name 		= c.PostForm("name")
		game_created.Category 		= c.PostForm("category")
		game_created.Brand 		= c.PostForm("brand")
		game_created.Year_released, _ 	= strconv.ParseInt(c.PostForm("year_released"), 10, 64)
		game_created.Price, _ 		= strconv.ParseFloat(c.PostForm("price"), 64)

		game_result := B_BUS_layer.CreateGameBUS(game_created)

		c.JSON(http.StatusOK, gin.H{"game_result": game_result})
	})
	// End create
	
	// Read details
	app.GET("/read-details/:id", func(c *gin.Context) {
		id := c.Param("id")
		game := B_BUS_layer.ReadDetailBUS(id)
		c.JSON(http.StatusOK, gin.H{"game_detail": game})
	})
	// End read details

	// Update game --> issue
	app.PUT("/update-game/:id", func(c *gin.Context) {
		game_updated := new(DTO.Game)

		game_updated.ID	 		= c.Param("id")
		game_updated.Name 		= c.PostForm("name")
		game_updated.Category 		= c.PostForm("category")
		game_updated.Brand 		= c.PostForm("brand")
		game_updated.Year_released, _ 	= strconv.ParseInt(c.PostForm("year_released"), 10, 64)
		game_updated.Price, _ 		= strconv.ParseFloat(c.PostForm("price"), 64)

		game_result := B_BUS_layer.UpdateGameBUS(game_updated)

		c.JSON(http.StatusOK, gin.H{"game_result": game_result})
	})
	// End update

	// Delete game
	app.DELETE("/delete-game/:id", func(c *gin.Context) {
		id := c.Param("id")
		game_result := B_BUS_layer.DeleteGameBUS(id)
		c.JSON(http.StatusOK, gin.H{"game_result": game_result})
	})
	// End delete

	// Search game
	app.GET("/search_game/:key_name", func(c *gin.Context) {
		search_name := c.Param("key_name")
		key_name := "%" + search_name + "%"
		games := B_BUS_layer.SearchGameBUS(key_name)
		c.JSON(http.StatusOK, gin.H{"game_list": games})
	})
	// End search

	// Run router
	app.Run()
	// --------------------------------------------------------------
	// --------------------------------------------------------------
	// END APP
}
