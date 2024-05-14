package base

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
)

// Options struct defines the configuration for the AppShell component.
type Opts struct {
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

// BuildOpts builds the options with the given opt
func BuildOpts(opts ...templwind.OptFunc[Opts]) *Opts {
	return templwind.BuildOpts(defaultOpts, opts...)
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

func WithTitle(title string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Title = title
	}
}

func WithMeta(meta templ.Component) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Meta = meta
	}
}

func WithFavicon(favicon string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Favicon = favicon
	}
}

func WithStylesheets(stylesheets ...string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Stylesheets = stylesheets
	}
}

func WithHeadScripts(headScripts ...string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.HeadScripts = headScripts
	}
}

func WithFooterScripts(footerScripts ...string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.FooterScripts = footerScripts
	}
}

func WithBodyCss(bodyCss string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.BodyCss = bodyCss
	}
}
