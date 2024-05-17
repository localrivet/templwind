package templwind

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Component interface with generic methods
type Component[T any] interface {
	New(opts ...OptFunc[T]) templ.Component
	NewWithOpt(opt *T) templ.Component
	WithOpts(opts ...OptFunc[T]) *T
}

// OptFunc is a generic function type for options
type OptFunc[T any] func(*T)

// New creates a new templ.Component with the given options
func New[T any](defaultOpts func() *T, tpl func(*T) templ.Component, opts ...OptFunc[T]) templ.Component {
	opt := WithOpts(defaultOpts, opts...)
	return tpl(opt)
}

// NewWithOpts creates a new templ.Component with the given opt
func NewWithOpts[T any](tpl func(*T) templ.Component, opts *T) templ.Component {
	return tpl(opts)
}

// WithOpts constructs the options with the given option functions
func WithOpts[T any](defaultOpts func() *T, opts ...OptFunc[T]) *T {
	opt := defaultOpts()
	for _, optFn := range opts {
		optFn(opt)
	}
	return opt
}

func Render(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)

	err := t.Render(context.Background(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}

	return nil
}
