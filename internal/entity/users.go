package entity

type Users struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	ReferencesId string `gorm:"<-:create;type:uuid" json:"references_id"`
	Username     string `json:"username" example:"john_doe"`
	Password     string `json:"password" example:"$2a$12$eixZaYVK1fsbw1ZfbX3OXe.PZyWJQ0Zf10hErsTQ6FVRHiA2vwLHu"` // Example of bcrypt-hashed password
}
