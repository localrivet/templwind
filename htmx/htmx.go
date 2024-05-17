package htmx

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func IsHtmxRequest(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}

func Redirect(c echo.Context, path string) error {
	if IsHtmxRequest(c) {
		c.Response().Header().Set("HX-Redirect", path)
		return c.NoContent(204)
	}

	return c.Redirect(302, path)
}

type LocationMap struct {
	Path   string `json:"path"`
	Target string `json:"target"`
}

func Location(c echo.Context, target LocationMap) error {
	if IsHtmxRequest(c) {
		data, err := json.Marshal(target)
		if err == nil {
			c.Response().Header().Set("HX-Location", string(data))
			return c.NoContent(204)
		}
	}

	return c.Redirect(302, target.Path)
}

func Trigger(c echo.Context, trigger string) {
	c.Response().Header().Set("HX-Trigger", trigger)
}
