package handle

import (
	"github.com/codeyhj/auth-admin/server/db"
	"github.com/codeyhj/auth-admin/server/modal"
	"github.com/codeyhj/auth-admin/server/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleToken(c echo.Context) error {
	d := db.GetDB()
	tModal := &[]modal.Token{}
	res := util.CreateRes()
	if err := d.Find(tModal).Error; err != nil {
		return c.JSON(http.StatusOK, res)
	}
	res.Data = tModal
	return c.JSON(http.StatusOK, res)
}
