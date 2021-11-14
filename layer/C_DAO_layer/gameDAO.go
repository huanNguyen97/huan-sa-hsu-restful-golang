package C_DAO_layer

import (
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

func CreateGameDAO(game_posted *DTO.Game) bool {
	db := database.Db_Connection()
	var game_result bool
	db.Table("game").Create(&game_posted)
	game_result = true
	return game_result
}

func ReadGameDetailDAO(id string) DTO.Game {
	db := database.Db_Connection()
	var game DTO.Game
	db.Table("game").Where("id = ?", id).Find(&game)
	return game
}

func UpdateGameDAO(game_posted *DTO.Game) bool {
	db := database.Db_Connection()
	var game_result bool
	db.Table("game").Save(&game_posted)
	game_result = true
	return game_result
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