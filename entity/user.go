package entity

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
}
