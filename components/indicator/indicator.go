package indicator

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
)

type Opts struct {
	IsUp  bool
	Value string
}

// New creates a new component
func New(opts ...templwind.OptFunc[Opts]) templ.Component {
	return templwind.New(defaultOpts, tpl, opts...)
}

// NewWithOpt creates a new component with the given opt
func NewWithOpts(opt *Opts) templ.Component {
	return templwind.NewWithOpts(tpl, opt)
}

// WithOpts builds the options with the given opt
func WithOpts(opts ...templwind.OptFunc[Opts]) *Opts {
	return templwind.WithOpts(defaultOpts, opts...)
}

func defaultOpts() *Opts {
	return &Opts{}
}

func IsUp(opts *Opts) {
	opts.IsUp = true
}

func Value(value string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Value = value
	}
}
