package card

import (
	"github.com/localrivet/templwind"
	"github.com/localrivet/templwind/components/indicator"
	"github.com/localrivet/templwind/components/link"

	"github.com/a-h/templ"
)

type Opts struct {
	ID            string
	Class         string
	LinkOpts      *link.Opts
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

// BuildOpts builds the options with the given opt
func BuildOpts(opts ...templwind.OptFunc[Opts]) *Opts {
	return templwind.BuildOpts(defaultOpts, opts...)
}

func defaultOpts() *Opts {
	return &Opts{
		Class: "bg-white border border-slate-200 rounded-lg shadow dark:bg-slate-800 dark:border-slate-700",
	}
}

func WithID(id string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.ID = id
	}
}

func WithClass(class string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Class = class
	}
}

func WithLinkOpts(linkOpts ...templwind.OptFunc[link.Opts]) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.LinkOpts = link.BuildOpts(linkOpts...)
	}
}

func WithTitle(title string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Title = title
	}
}

func WithSubTitle(subTitle string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.SubTitle = subTitle
	}
}

func WithLead(lead string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Lead = lead
	}
}

func WithHeadIndicator(headIndicator *indicator.Opts) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HeadIndicator = headIndicator
	}
}

func WithComponents(components ...templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Components = components
	}
}

func WithButtons(buttons templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Buttons = buttons
	}
}
