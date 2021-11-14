package DTO

type Game struct {
	ID 				string 	`json:"id"`			
	Name 			string 	`json:"name"`
	Category 		string 	`json:"category"`
	Brand 			string 	`json:"brand"`
	Year_released 	int64 	`json:"year_released"`
	Price 			float64 `json:"price"`
}
