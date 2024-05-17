package index

import (
	"github.com/labstack/echo/v4"
	"github.com/localrivet/templwind/docs/internal/svc"
)

func Routes(svcCtx *svc.ServiceContext, parentGroup *echo.Group) {

	// index
	g := parentGroup.Group("")

	// base route
	g.GET("", NewController(svcCtx).Index)
}
