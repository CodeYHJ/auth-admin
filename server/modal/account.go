package modal

import (
	"time"
)

type Account struct {
	ID        uint      `json:"id" gorm:"column:id; primaryKey; autoIncrement; not null; type:uint;comment:id"`
	UserName  string    `json:"username" gorm:"column:username; not null; type:varchar(500); comment:user name"`
	Paw       string    `json:"paw" gorm:"column:paw; not null; type:varchar(500); comment:password"`
	CreatedAt time.Time `gorm:"column:created_at; type:time" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at; type:time" json:"updated_at"`
}

func (Account) TableName() string {
	return "accounts"
}
