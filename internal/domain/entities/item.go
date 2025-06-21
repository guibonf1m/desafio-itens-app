package entities

type Item struct {
	ID          int     `json:"id"`
	Code        string  `json:"code"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"quantidade"`
	Status      string  `json:"status"`
	Created_at  int     `json:"created"`
	Update_at   int     `json:"update"`
}
