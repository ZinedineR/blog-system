package entity

import "time"

type Posts struct {
	Id               int        `json:"-" gorm:"primaryKey;autoIncrement"`
	ReferencesId     string     `gorm:"<-:create;unique;type:uuid" json:"references_id"`
	UserReferencesId string     `json:"user_references_id" gorm:"not null"`
	Title            string     `json:"title" gorm:"type:varchar(255);not null"`
	Content          string     `json:"content" gorm:"type:text;not null"`
	Tags             string     `json:"tags" gorm:"type:varchar(255)"`
	Published        bool       `json:"published"`
	LastPublishedAt  *time.Time `json:"last_published_at"`
	CreatedAt        time.Time  `gorm:"autoCreateTime;<-:create" json:"created_at"`
	UpdatedAt        *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	//relationship
	Comments []*Comments `json:"comments,omitempty" gorm:"foreignKey:PostReferencesId;references:ReferencesId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Author   *Users      `json:"author,omitempty" gorm:"foreignKey:UserReferencesId;references:ReferencesId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
