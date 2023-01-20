package model

type User struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Balance      int           `json:"Balance"`
	Transections []Transection `json:"transections"`
}

type Transection struct {
	SenderID   string `json:"sender"`
	RecieverID string `json:"reciver"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
}
