package entity

type CollectionEntity struct {
	Id     int    `json:"id" db:"id"`
	UserId string `json:"userId" db:"user_id"`
	NavId  string `json:"navId" db:"nav_id"`
}
