package modal

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type JSON json.RawMessage

func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}
func (Client) TableName() string {
	return "oauth2_clients"
}

type Client struct {
	ID        string    `gorm:"column:id; primaryKey; not null; type:varchar(500);comment:id" json:"id"`
	Secret    string    `gorm:"column:secret; not null; type:varchar(500); comment:secret" json:"secret"`
	Domain    string    `gorm:"column:domain; not null; type:varchar(500); comment:domain" json:"domain"`
	Data      JSON      `gorm:"column:data; type:json" json:"data"`
	UpdatedAt time.Time `gorm:"column:updated_at; type:time" json:"updated_at"`
}
