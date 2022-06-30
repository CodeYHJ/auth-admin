package handle

import (
	"encoding/json"
	"github.com/codeyhj/auth-admin/server/db"
	"github.com/codeyhj/auth-admin/server/modal"
	"github.com/codeyhj/auth-admin/server/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func HandleClient(c echo.Context) error {
	d := db.GetDB()
	cModal := &[]modal.Client{}
	res := util.CreateRes()
	if err := d.Find(cModal).Error; err != nil {
		return c.JSON(http.StatusOK, res)
	}
	res.Data = cModal
	return c.JSON(http.StatusOK, res)
}

func HandleClientAdd(c echo.Context) error {
	q := &modal.ClientAddQuery{}
	res := util.CreateRes()
	if err := c.Bind(q); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(q); err != nil {
		return c.JSON(http.StatusOK, res)
	}
	domain := q.Domain
	secret := q.Secret
	id := q.ID
	d := db.GetDB()
	cModal := &modal.Client{ID: id}
	d.First(&cModal)
	if cModal.Domain == "" {
		updateAt := time.Now()
		cModal.Secret = secret
		cModal.Domain = domain
		cModal.UpdatedAt = updateAt
		jsonData, err := json.Marshal(&cModal)
		if err != nil {
			return c.JSON(http.StatusOK, res)
		}
		cModal.Data = jsonData
		if err := d.Create(cModal).Error; err != nil {
			return c.JSON(http.StatusOK, res)
		}
	}

	return c.JSON(http.StatusOK, res)
}

func HandleClientEdit(c echo.Context) error {
	q := &modal.ClientEditQuery{}
	res := util.CreateRes()
	if err := c.Bind(q); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(q); err != nil {
		return c.JSON(http.StatusOK, res)
	}
	domain := q.Domain
	secret := q.Secret
	id := q.ID
	d := db.GetDB()
	cModal := &modal.Client{ID: id}
	d.First(&cModal)
	if cModal.Domain != "" {
		updateAt := time.Now()
		jsonData, err := json.Marshal(&modal.Client{Domain: domain, Secret: secret, ID: id, UpdatedAt: updateAt})
		if err != nil {
			return c.JSON(http.StatusOK, res)
		}
		cModal.Data = jsonData
		d.Model(&modal.Client{ID: id}).Updates(&modal.Client{Domain: domain, Secret: secret, Data: jsonData, UpdatedAt: updateAt})
	}
	return c.JSON(http.StatusOK, res)

}
