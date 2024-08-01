package model

import (
	"time"
	//"github.com/lib/pq"
)

type BlogStatus int
const (
	Draft BlogStatus = iota
	Published
	Closed
	Active
	Famous
)

type BlogCategory int
const (
	Destinations BlogCategory = iota
	Travelogues
	Activities
	Gastronomy
	Tips
	Culture
	Accomodation
)

type Blog struct {
	Id				int				`json:"id"`
	UserId			int 			`json:"userId"`
	Title 			string			`json:"title"`
	Description		string			`json:"description"`
	CreationTime 	time.Time 		`json:"creationTime"`
	Image			string			`json:"image"`
	Status 			BlogStatus		`json:"status"`
	Category		BlogCategory	`json:"category"`
}