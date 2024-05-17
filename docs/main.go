package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/localrivet/templwind/docs/internal/config"
	"github.com/localrivet/templwind/docs/internal/svc"
	"github.com/localrivet/templwind/docs/modules/docs"
	"github.com/localrivet/templwind/docs/modules/frontend"
)

//go:embed etc/config.yaml
var configFile embed.FS

//go:embed all:assets
var assetsPath embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configBytes, err := configFile.ReadFile("etc/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// Expand environment variables
	configBytes = []byte(os.ExpandEnv(string(configBytes)))

	var c config.Config
	err = config.LoadConfigFromYamlBytes(configBytes, &c)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create the service context with a new PocketBase instance
	svcCtx := svc.NewServiceContext(&c)

	// Create a new Echo instance

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

	// Register the routes
	docs.Routes(svcCtx, e.Group("/docs"))
	frontend.Routes(svcCtx, e.Group(""))

	// Start the server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", c.Port)))
}
