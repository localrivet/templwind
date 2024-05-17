package frontend

import (
	"github.com/labstack/echo/v4"
	"github.com/localrivet/templwind/docs/internal/svc"
	"github.com/localrivet/templwind/docs/modules/frontend/index"
)

func Routes(svcCtx *svc.ServiceContext, g *echo.Group) {
	// index
	index.Routes(svcCtx, g)
}
