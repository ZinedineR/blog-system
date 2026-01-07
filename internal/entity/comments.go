package entity

import "time"

type Comments struct {
	Id               int        `json:"-" gorm:"primaryKey;autoIncrement"`
	ReferencesId     string     `gorm:"<-:create;unique;type:uuid" json:"references_id"`
	PostReferencesId string     `json:"post_references_id" gorm:"not null"`
	UserReferencesId *string    `json:"user_references_id"`
	Content          string     `json:"content" gorm:"type:text;not null"`
	CreatedAt        time.Time  `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt        *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	//relationship
	Author *Users `json:"author,omitempty" gorm:"foreignKey:UserReferencesId;references:ReferencesId;constraint:OnDelete:SET NULL"`
}
