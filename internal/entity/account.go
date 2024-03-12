package entity

type Account struct {
	// user id
	UserId int64
	// user phone number
	PhoneNumber string
	// username like @timohahaa (without @ symbol)
	Username string
	// session as a string
	SessionString string
}
