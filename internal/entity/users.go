package entity

type Users struct {
	Id           int    `json:"-" gorm:"primaryKey;autoIncrement"`
	ReferencesId string `gorm:"<-:create;unique;type:uuid" json:"references_id"`
	Username     string `json:"username" example:"john_doe"`
	Password     string `json:"-"`
}
