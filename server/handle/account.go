package handle

import (
	"github.com/codeyhj/auth-admin/server/db"
	"github.com/codeyhj/auth-admin/server/modal"
	"github.com/codeyhj/auth-admin/server/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func HandleAccount(c echo.Context) error {
	d := db.GetDB()
	aModal := &[]modal.Account{}
	res := util.CreateRes()
	if err := d.Find(aModal).Error; err != nil {
		return c.JSON(http.StatusOK, res)
	}
	res.Data = aModal
	return c.JSON(http.StatusOK, res)
}

func HandleAccountAdd(c echo.Context) error {
	q := &modal.AccountAddQuery{}
	res := util.CreateRes()
	if err := c.Bind(q); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(q); err != nil {
		return c.JSON(http.StatusOK, res)
	}
	u := q.UserName
	p := q.Paw
	d := db.GetDB()
	cModal := &modal.Account{UserName: u}
	d.First(&cModal)
	if cModal.ID == 0 {
		updateAt := time.Now()
		cModal.Paw = p
		cModal.UpdatedAt = updateAt
		cModal.CreatedAt = updateAt
		if err := d.Create(cModal).Error; err != nil {
			return c.JSON(http.StatusOK, res)
		}
	}

	return c.JSON(http.StatusOK, res)
}

func HandleAccountEdit(c echo.Context) error {
	q := &modal.AccountEditQuery{}
	res := util.CreateRes()
	if err := c.Bind(q); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(q); err != nil {
		return c.JSON(http.StatusOK, res)
	}
	id := q.ID
	p := q.Paw
	d := db.GetDB()
	cModal := &modal.Account{ID: id}
	d.First(&cModal)
	if cModal.ID != 0 {
		updateAt := time.Now()
		cModal.Paw = p
		cModal.UpdatedAt = updateAt
		d.Model(&modal.Account{ID: id}).Updates(&modal.Account{Paw: p, UpdatedAt: updateAt})
	}

	return c.JSON(http.StatusOK, res)
}
