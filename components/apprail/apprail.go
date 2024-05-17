package apprail

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
	"github.com/localrivet/templwind/components/link"
)

// Opts defines the options for the AppRail component
type Opts struct {
	Background  string
	Border      string
	Width       string
	Height      string
	Gap         string
	Hover       string
	Active      string
	Spacing     string
	AspectRatio string
	Lead        templ.Component
	Trail       templ.Component
	MenuItems   []*link.Opts
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

// DefaultOpts provides the default options for the AppRail component
func defaultOpts() *Opts {
	return &Opts{
		Background:  "bg-surface-100-800-token",
		Border:      "",
		Width:       "w-20",
		Height:      "h-full",
		Gap:         "gap-0",
		Hover:       "bg-primary-hover-token",
		Active:      "bg-primary-active-token",
		Spacing:     "space-y-1",
		AspectRatio: "aspect-square",
		Lead:        nil,
		Trail:       nil,
		MenuItems:   make([]*link.Opts, 0),
	}
}

func Background(background string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Background = background
	}
}

func Border(border string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Border = border
	}
}

func Width(width string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Width = width
	}
}

func Height(height string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Height = height
	}
}

func Gap(gap string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Gap = gap
	}
}

func Hover(hover string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Hover = hover
	}
}

func Active(active string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Active = active
	}
}

func Spacing(spacing string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Spacing = spacing
	}
}

func AspectRatio(aspectRatio string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.AspectRatio = aspectRatio
	}
}

func Lead(lead templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Lead = lead
	}
}

func MenuItems(menuItems ...*link.Opts) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.MenuItems = append(opts.MenuItems, menuItems...)
	}
}
