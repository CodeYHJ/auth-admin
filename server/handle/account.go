package handle

import (
	"github.com/codeyhj/auth-admin/server/db"
	"github.com/codeyhj/auth-admin/server/modal"
	"github.com/codeyhj/auth-admin/server/util"
	"github.com/labstack/echo/v4"
	"net/http"
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
