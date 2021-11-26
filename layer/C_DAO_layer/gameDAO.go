package C_DAO_layer

import (
	// Go imported primary
	"reflect"

	// Go import go module
	"github.com/huanNguyen97/huan-sa-hsu-golang/database"
	"github.com/huanNguyen97/huan-sa-hsu-golang/DTO"
)

func ReadGamesDAO() []DTO.Game {
	db := database.Db_Connection()
	var game []DTO.Game
	db.Table("game").Find(&game)
	return game
}

func CreateGameDAO(game_created *DTO.Game) bool {
	// Check condition type of data about number type
	checkTypeOfYear := reflect.TypeOf(game_created.Year_released)
	checkTypeOfPrice := reflect.TypeOf(game_created.Price)

	if checkTypeOfYear.String() != "int64" || 
	   checkTypeOfPrice.String() != "float64" {
		return false
	} else {
		db := database.Db_Connection()

		game_result := false
		db.Table("game").Create(&game_created)
		game_result = true
	
		return game_result
	}
}

func ReadGameDetailDAO(id string) DTO.Game {
	db := database.Db_Connection()
	var game DTO.Game
	db.Table("game").Where("id = ?", id).Find(&game)
	return game
}

func UpdateGameDAO(game_updated *DTO.Game) bool {
	checkTypeOfYear := reflect.TypeOf(game_updated.Year_released)
	checkTypeOfPrice := reflect.TypeOf(game_updated.Price)

	if checkTypeOfYear.String() != "int64" || 
	   checkTypeOfPrice.String() != "float64" {
		return false 
	} else {
		db := database.Db_Connection()
		db.Table("game").Save(&game_updated)
		return true
	}
}

func DeleteGameDAO(id string) bool {
	db := database.Db_Connection()
	var game DTO.Game
	db.Table("game").Where("id = ?", id).Delete(&game)
	game_result := true 
	return game_result
}

func SearchGameDAO(key_name string) []DTO.Game {
	db := database.Db_Connection()
	var games []DTO.Game
	db.Table("game").Where("name LIKE ?", key_name).Find(&games)
	return games
}