package main

import (
	// Go importing pkg primary
	// "fmt"
	"net/http"
	"strconv"

	// Go importing module
	"github.com/gin-gonic/gin"
	
	// Go importing owner
	"github.com/huanNguyen97/huan-sa-hsu-golang/layer/B_BUS_layer"
	"github.com/huanNguyen97/huan-sa-hsu-golang/DTO"
)

func main() {
	// APP CREATED 
	// --------------------------------------------------------------
	// --------------------------------------------------------------
	app := gin.Default()

	app.LoadHTMLGlob("layer/A_GUI_layer/*")
	app.Static("/css", "./css")

	// Read games
	app.GET("/", func(c *gin.Context) {
		games := B_BUS_layer.ReadGamesBUS()
		c.HTML(http.StatusOK, "read_games.html", gin.H{"games": games})
	})
	// End read
	
	// Create game
	app.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create_game.html", gin.H{"create": "create"})
	})

	app.POST("/create", func(c *gin.Context) {
		Name 			:= c.PostForm("name")
		Category 		:= c.PostForm("category")
		Brand 			:= c.PostForm("brand")
		Year_released 	:= c.PostForm("year_released")
		Price 			:= c.PostForm("price")

		game_posted 					:= new(DTO.Game)
		game_posted.ID 					= ""
		game_posted.Name 				= Name
		game_posted.Category 			= Category
		game_posted.Brand 				= Brand
		game_posted.Year_released, _ 	= strconv.ParseInt(Year_released, 10, 64)
		game_posted.Price, _ 			= strconv.ParseFloat(Price, 64)

		B_BUS_layer.CreateGameBUS(game_posted)
		games := B_BUS_layer.ReadGamesBUS()

		c.HTML(http.StatusOK, "read_games.html", gin.H{"games": games})
	})
	// End create
	
	// Read details
	app.GET("/read_details/:id", func(c *gin.Context) {
		id := c.Param("id")
		game := B_BUS_layer.ReadDetailBUS(id)
		c.HTML(http.StatusOK, "read_game_details.html", gin.H{"game": game})
	})
	// End read details

	// Update game
	app.GET("/update_game/:id", func(c *gin.Context) {
		id := c.Param("id")
		game := B_BUS_layer.ReadDetailBUS(id)
		c.HTML(http.StatusOK, "update_game.html", gin.H{"game": game})
	})

	app.POST("/update_game/:id", func(c *gin.Context) {
		ID				:= c.PostForm("id")
		Name 			:= c.PostForm("name")
		Category 		:= c.PostForm("category")
		Brand 			:= c.PostForm("brand")
		Year_released 	:= c.PostForm("year_released")
		Price 			:= c.PostForm("price")

		game_posted 					:= new(DTO.Game)
		game_posted.ID 					= ID
		game_posted.Name 				= Name
		game_posted.Category 			= Category
		game_posted.Brand 				= Brand
		game_posted.Year_released, _ 	= strconv.ParseInt(Year_released, 10, 64)
		game_posted.Price, _ 			= strconv.ParseFloat(Price, 64)

		B_BUS_layer.UpdateGameBUS(game_posted)

		games := B_BUS_layer.ReadGamesBUS()
		c.HTML(http.StatusOK, "read_games.html", gin.H{"games": games})
	})
	// End update

	// Delete game
	app.POST("/delete_game/:id", func(c *gin.Context) {
		id := c.Param("id")
		B_BUS_layer.DeleteGameBUS(id)
		games := B_BUS_layer.ReadGamesBUS()
		c.HTML(http.StatusOK, "read_games.html", gin.H{"games": games})
	})
	// End delete

	// Search game
	app.POST("/search_game", func(c *gin.Context) {
		search_name := c.PostForm("text_of_searching")
		key_name := "%" + search_name + "%"
		games := B_BUS_layer.SearchGameBUS(key_name)
		c.HTML(http.StatusOK, "read_games.html", gin.H{"games": games})
	})
	// End search

	// Run router
	app.Run("localhost:8000")
	// --------------------------------------------------------------
	// --------------------------------------------------------------
	// END APP
}