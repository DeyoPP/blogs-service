package model

// import (
// 	"time"
// 	//"github.com/lib/pq"
// )

// Uncomment the following if you intend to use them
// type BlogStatus int
// const (
// 	Draft BlogStatus = iota
// 	Published
// 	Closed
// 	Active
// 	Famous
// )

// type BlogCategory int
// const (
// 	Destinations BlogCategory = iota
// 	Travelogues
// 	Activities
// 	Gastronomy
// 	Tips
// 	Culture
// 	Accomodation
// )

type Blog struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`          // Primary key with auto-increment
	Name        string    `gorm:"column:name;type:varchar(255)" json:"name"`   // Column name and type
	Description string    `gorm:"column:description;type:text" json:"description"` // Column name and type
	// CreationTime time.Time `json:"creationTime"`
	// Image        string    `json:"image"`
	// Status       BlogStatus `json:"status"`
	// Category     BlogCategory `json:"category"`
}
