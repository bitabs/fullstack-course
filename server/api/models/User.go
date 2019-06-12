package models

type User struct {
	ID			ID			`json:"id"`
	Name 		Name		`json:"name"`
	Gender		string		`json:"gender"`
	Location	Location	`json:"location"`
	Email		string		`json:"email"`
	Phone		string		`json:"phone"`
	Login		Login		`json:"login"`
	DOB			DOB			`json:"dob"`
	Registered	Registered	`json:"registered"`
	Cell		string		`json:"cell"`
	Picture		Picture		`json:"picture"`
	NAT			string		`json:"nat"`
}

type ID struct {
	Name		string		`json:"name"`
	Value		string		`json:"value"`
}

type Name struct {
	Title		string		`json:"title"`
	First		string		`json:"first"`
	Last		string		`json:"last"`
}

type Location struct {
	Street		string		`json:"street"`
	City		string		`json:"city"`
	State		string		`json:"state"`
	Postcode	string		`json:"postcode"`
	Coordinates	Coordinates	`json:"coordinates"`
	Timezone	TimeZone	`json:"timezone"`
}

type Coordinates struct {
	Latitude	string		`json:"latitude"`
	Longitude	string		`json:"longitude"`
}

type TimeZone struct {
	Offset		string		`json:"offset"`
	Description	string		`json:"description"`
}

type Login struct {
	UUID		string		`json:"uuid"`
	Username	string		`json:"username"`
	Password	string		`json:"password"`
	Salt		string		`json:"salt"`
	MD5			string		`json:"md5"`
	SHA1		string		`json:"sha1"`
	SHA256		string		`json:"sha256"`
}

type DOB struct {
	Date		string		`json:"date"`
	Age			int			`json:"age"`
}

type Registered struct {
	Date		string		`json:"date"`
	Age			int			`json:"age"`
}

type Picture struct {
	Large		string		`json:"large"`
	Medium		string		`json:"medium"`
	Thumbnail	string		`json:"thumbnail"`
}
