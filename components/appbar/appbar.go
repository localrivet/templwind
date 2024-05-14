package appbar

import (
	"github.com/a-h/templ"
	"github.com/localrivet/templwind"
)

// Opts defines the options for the AppBar component
type Opts struct {
	Background        string
	Border            string
	Padding           string
	Shadow            string
	Spacing           string
	GridColumns       string
	Gap               string
	RegionRowMain     string
	RegionRowHeadline string
	Lead              templ.Component
	Trail             templ.Component
	Label             string
	Labelledby        string
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

// DefaultOpts provides the default options for the AppBar component
func defaultOpts() *Opts {
	return &Opts{
		Background:        "bg-surface-100-800-token",
		Border:            "",
		Padding:           "p-4",
		Shadow:            "",
		Spacing:           "space-y-4",
		GridColumns:       "grid-cols-[auto_1fr_auto]",
		Gap:               "gap-4",
		RegionRowMain:     "",
		RegionRowHeadline: "",
		Label:             "",
		Labelledby:        "",
	}
}

func WithBackground(background string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Background = background
	}
}

func WithBorder(border string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Border = border
	}
}

func WithPadding(padding string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Padding = padding
	}
}

func WithShadow(shadow string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Shadow = shadow
	}
}

func WithSpacing(spacing string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Spacing = spacing
	}
}

func WithGridColumns(gridColumns string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.GridColumns = gridColumns
	}
}

func WithGap(gap string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.Gap = gap
	}
}

func WithRegionRowMain(regionRowMain string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.RegionRowMain = regionRowMain
	}
}

func WithRegionRowHeadline(regionRowHeadline string) templwind.OptFunc[Opts] {
	return func(o *Opts) {
		o.RegionRowHeadline = regionRowHeadline
	}
}
