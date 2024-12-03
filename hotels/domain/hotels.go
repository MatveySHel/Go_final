package domain

type Hotel struct {
	ID 		int			`db:"id"`
	Name 	string		`db:"name"`
	Price 	int			`db:"price"`
	City 	string		`db:"city"`
}