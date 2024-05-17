package sidenav

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
	"github.com/localrivet/templwind/components/link"
)

type Opts struct {
	ID             string
	ContainerClass string
	Submenu        []*link.Opts
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

// DefaultOpts provides the default options for the AppBar component
func defaultOpts() *Opts {
	return &Opts{
		ID:      "sidenav",
		Submenu: []*link.Opts{},
	}
}

func ID(id string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.ID = id
	}
}

func ContainerClass(class string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.ContainerClass = class
	}
}

func Submenu(submenu ...*link.Opts) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Submenu = append(opts.Submenu, submenu...)
	}
}
