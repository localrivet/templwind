package index

import (
	"github.com/labstack/echo/v4"
	"github.com/localrivet/templwind"
	"github.com/localrivet/templwind/docs/internal/svc"
	"github.com/localrivet/templwind/docs/modules/docs/index/views"
)

type Controller struct {
	svcCtx *svc.ServiceContext
}

func NewController(svcCtx *svc.ServiceContext) *Controller {
	return &Controller{
		svcCtx: svcCtx,
	}
}

func (c *Controller) Index(ctx echo.Context) error {
	section := ctx.Param("section")
	if section == "" {
		section = "docs"
	}

	// subSection := ctx.Param("subSection")

	return templwind.Render(ctx, 200, views.Index(ctx, section))
}

func (c *Controller) LoadMDContent(ctx echo.Context) error {
	// return templwind.Render(ctx, 404, views.NotFound(ctx))
	return nil
}
