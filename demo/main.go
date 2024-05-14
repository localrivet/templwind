package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/localrivet/templwind/demo/app"
	"github.com/localrivet/templwind/demo/app2"
)

//go:embed all:assets
var assetsPath embed.FS

func main() {
	e := echo.New()
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	// Register the static file server
	subFS, err := fs.Sub(assetsPath, "assets")
	if err != nil {
		log.Fatalf("Failed to create sub filesystem: %v", err)
	}
	e.StaticFS("/static/", subFS)

	// Register the index handler
	e.GET("/", func(ctx echo.Context) error {
		return render(ctx, http.StatusOK, app.Index())
	})

	e.GET("/app2", func(ctx echo.Context) error {
		return render(ctx, http.StatusOK, app2.Index())
	})

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}

func render(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)

	err := t.Render(context.Background(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}

	return nil
}
