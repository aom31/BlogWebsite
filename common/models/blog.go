package models

type Blog struct {
	ID     uint   `json: "id"`
	Title  string `json: "title"`
	Desc   string `json:"desc" description: "description"`
	Image  string `json:"image"`
	UserID string `json:"user_id"`
	User   User   `json:"user"; gorm:"foreignkey:UserID"` //foreignkey is key data external model
}
