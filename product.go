package iiiproject

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	LastUpdate string `json:"last_update"`
}
