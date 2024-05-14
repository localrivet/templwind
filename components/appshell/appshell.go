package appshell

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
)

// Options struct defines the configuration for the AppShell component.
type Opts struct {
	ScrollbarGutter string
	RegionPage      string
	Header          templ.Component
	SidebarLeft     templ.Component
	SidebarRight    templ.Component
	PageHeader      templ.Component
	PageFooter      templ.Component
	Footer          templ.Component
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
		ScrollbarGutter: "auto",
		Header:          nil,
		SidebarLeft:     nil,
		SidebarRight:    nil,
	}
}

func WithScrollbarGutter(scrollbarGutter string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.ScrollbarGutter = scrollbarGutter
	}
}

func WithRegionPage(regionPage string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.RegionPage = regionPage
	}
}

func WithHeader(header templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Header = header
	}
}

func WithSidebarLeft(sidebarLeft templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.SidebarLeft = sidebarLeft
	}
}

func WithSidebarRight(sidebarRight templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.SidebarRight = sidebarRight
	}
}

func WithPageHeader(pageHeader templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.PageHeader = pageHeader
	}
}

func WithPageFooter(pageFooter templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.PageFooter = pageFooter
	}
}

func WithFooter(footer templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Footer = footer
	}
}
