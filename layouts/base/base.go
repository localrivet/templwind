package base

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/localrivet/templwind"
)

// Options struct defines the configuration for the AppShell component.
type Opts struct {
	EchoCtx       echo.Context
	Title         string
	Meta          templ.Component
	Favicon       string
	Stylesheets   []string
	HeadScripts   []string
	FooterScripts []string
	BodyCss       string
}

// New creates a new component
func New(opts ...templwind.OptFunc[Opts]) templ.Component {
	return templwind.New(defaultOpts, tpl, opts...)
}

// NewWithOpts creates a new component with the given opt
func NewWithOpts(opt *Opts) templ.Component {
	return templwind.NewWithOpts(tpl, opt)
}

// WithOpts builds the options with the given opt
func WithOpts(opts ...templwind.OptFunc[Opts]) *Opts {
	return templwind.WithOpts(defaultOpts, opts...)
}

// Default options for the AppShell component.
func defaultOpts() *Opts {
	return &Opts{
		Title:         "TemplWind",
		Meta:          nil,
		Favicon:       "",
		Stylesheets:   []string{},
		HeadScripts:   []string{},
		FooterScripts: []string{},
		BodyCss:       "",
	}
}

func EchoCtx(e echo.Context) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.EchoCtx = e
	}
}

func Title(title string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Title = title
	}
}

func Meta(meta templ.Component) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Meta = meta
	}
}

func Favicon(favicon string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Favicon = favicon
	}
}

func Stylesheets(stylesheets ...string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Stylesheets = stylesheets
	}
}

func HeadScripts(headScripts ...string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.HeadScripts = headScripts
	}
}

func FooterScripts(footerScripts ...string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.FooterScripts = footerScripts
	}
}

func BodyCss(bodyCss string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.BodyCss = bodyCss
	}
}
