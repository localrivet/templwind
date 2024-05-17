package appshell

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
)

// Options struct defines the configuration for the AppShell component.
type Opts struct {
	Header              templ.Component
	HeaderClasses       string
	SidebarLeft         templ.Component
	SidebarLeftClasses  string
	SidebarRight        templ.Component
	SidebarRightClasses string
	PageHeader          templ.Component
	PageHeaderClasses   string
	PageFooter          templ.Component
	PageFooterClasses   string
	Footer              templ.Component
	FooterClasses       string
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
		Header:       nil,
		SidebarLeft:  nil,
		SidebarRight: nil,
	}
}

func Header(header templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Header = header
	}
}

func HeaderClasses(headerClasses string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HeaderClasses = headerClasses
	}
}

func SidebarLeft(sidebarLeft templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.SidebarLeft = sidebarLeft
	}
}

func SidebarLeftClasses(sidebarLeftClasses string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.SidebarLeftClasses = sidebarLeftClasses
	}
}

func SidebarRight(sidebarRight templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.SidebarRight = sidebarRight
	}
}

func SidebarRightClasses(sidebarRightClasses string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.SidebarRightClasses = sidebarRightClasses
	}
}

func PageHeader(pageHeader templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.PageHeader = pageHeader
	}
}

func PageHeaderClasses(pageHeaderClasses string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.PageHeaderClasses = pageHeaderClasses
	}
}

func PageFooter(pageFooter templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.PageFooter = pageFooter
	}
}

func PageFooterClasses(pageFooterClasses string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.PageFooterClasses = pageFooterClasses
	}
}

func Footer(footer templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Footer = footer
	}
}

func FooterClasses(footerClasses string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.FooterClasses = footerClasses
	}
}
