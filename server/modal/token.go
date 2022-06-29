package modal

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	gorm.Model
	ID        string    `gorm:"column:id; primaryKey; autoIncrement; not null; type:uint;comment:id" json:"id"`
	Code      string    `gorm:"column:code; not null; type:varchar(500); comment:code" json:"code"`
	Access    string    `gorm:"column:access; not null; type:varchar(500); comment:access" json:"access"`
	Refresh   string    `gorm:"column:refresh; not null; type:varchar(500); comment:refresh" json:"refresh"`
	Data      JSON      `gorm:"column:data; type:json" json:"data"`
	ExpiresAt time.Time `gorm:"column:expires_at; column:expires_at; not null; type:time;" json:"expiresAt"`
}

func (Token) TableName() string {
	return "oauth2_tokens"
}
