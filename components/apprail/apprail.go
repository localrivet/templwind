package apprail

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
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
	}
}

func WithBackground(background string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Background = background
	}
}

func WithBorder(border string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Border = border
	}
}

func WithWidth(width string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Width = width
	}
}

func WithHeight(height string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Height = height
	}
}

func WithGap(gap string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Gap = gap
	}
}

func WithHover(hover string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Hover = hover
	}
}

func WithActive(active string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Active = active
	}
}

func WithSpacing(spacing string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Spacing = spacing
	}
}

func WithAspectRatio(aspectRatio string) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.AspectRatio = aspectRatio
	}
}

func WithLead(lead templ.Component) templwind.OptFunc[Opts] {
	return func(opts *Opts) {
		opts.Lead = lead
	}
}
