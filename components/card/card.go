package card

import (
	"github.com/localrivet/templwind"
	"github.com/localrivet/templwind/components/indicator"

	"github.com/a-h/templ"
)

type Opts struct {
	ID            string
	Class         string
	Title         string
	SubTitle      string
	Lead          string
	HeadIndicator *indicator.Opts
	Components    []templ.Component
	Buttons       templ.Component
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

func defaultOpts() *Opts {
	return &Opts{
		Class: "bg-white border border-slate-200 rounded-lg shadow dark:bg-slate-800 dark:border-slate-700",
	}
}

func ID(id string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.ID = id
	}
}

func Class(class string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Class = class
	}
}

func Title(title string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Title = title
	}
}

func SubTitle(subTitle string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.SubTitle = subTitle
	}
}

func Lead(lead string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Lead = lead
	}
}

func HeadIndicator(headIndicator *indicator.Opts) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HeadIndicator = headIndicator
	}
}

func Components(components ...templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Components = components
	}
}

func Buttons(buttons templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Buttons = buttons
	}
}
