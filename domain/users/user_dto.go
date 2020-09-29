package users

type User struct {
	ID string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	FullName string `json:"full_name"`
	Age int `json:"age"`
	EmailId string `json:"email_id"`
	Address Address `json:"address"`
	Picture Picture `json:"picture"`
}

type Address struct {
	StreetNumber string `json:"street_number"`
	StreetName string `json:"street_name"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type Picture struct {
	ID int64 `json:"id"`
	Url string `json:"url"`
}