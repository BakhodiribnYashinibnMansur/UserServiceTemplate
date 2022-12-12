package model

type UserData struct {
	Id        int64  `db:"id"`
	FirstName string `db:"firstname"`
	LastName  string `db:"lastname"`
	Phone     string `db:"phone"`
	Password  string `db:"password"`
}
