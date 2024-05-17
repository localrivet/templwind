package appbar

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
)

type Class interface {
	Border() string
	BorderTop() string
	BorderBottom() string
	BorderLeft() string
	BorderRight() string
}

// Opts defines the options for the AppBar component
type Opts struct {
	AppBarClasses   string
	Lead            templ.Component
	LeadClasses     string
	Trail           templ.Component
	TrailClasses    string
	Headline        templ.Component
	HeadlineClasses string
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
		AppBarClasses:   "app-bar flex flex-col bg-surface-100-800-token space-y-4 p-4 w-full",
		LeadClasses:     "app-bar-slot-lead flex-none flex justify-between items-center",
		TrailClasses:    "app-bar-slot-trail flex-none flex items-center space-x-4",
		HeadlineClasses: "app-bar-row-headline",
	}
}

func AppBarClasses(appBarClasses string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.AppBarClasses = appBarClasses
	}
}

func Lead(lead templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Lead = lead
	}
}

func LeadClasses(leadClasses string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.LeadClasses = leadClasses
	}
}

func Trail(trail templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Trail = trail
	}
}

func TrailClasses(trailClasses string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.TrailClasses = trailClasses
	}
}

func Headline(headline templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Headline = headline
	}
}

func HeadlineClasses(headlineClasses string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.HeadlineClasses = headlineClasses
	}
}
