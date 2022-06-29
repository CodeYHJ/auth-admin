package modal

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	ID       uint      `json:"id" gorm:"column:id; primaryKey; autoIncrement; not null; type:uint;comment:id"`
	UserName string    `json:"username" gorm:"column:username; not null; type:varchar(500); comment:user name"`
	Paw      string    `json:"paw" gorm:"column:paw; not null; type:varchar(500); comment:password"`
	CreateAt time.Time `gorm:"column:create_at; type:time" json:"create_at"`
}

func (Account) TableName() string {
	return "accounts"
}
